package main

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"dagger/ccp/internal/dagger"
)

const (
	mcpDBName = "MCP"
	ssDBName  = "SS"
	dumpsDir  = "hack/ccp/integration/data"
)

// DatabaseExecutionMode defines the different modes in which the e2e tests can
// operate with the application databases.
type DatabaseExecutionMode string

const (
	// UseCached is the default mode that relies on whatever is the existing
	// MySQL service state.
	UseCached DatabaseExecutionMode = "USE_CACHED"
	// UseDumps attempts to configure the MySQL service using the database dumps
	// previously generated.
	UseDumps DatabaseExecutionMode = "USE_DUMPS"
	// ForceDrop drops the existing databases forcing the application to
	// recreate them using Django migrations.
	ForceDrop DatabaseExecutionMode = "FORCE_DROP"
)

func (m *CCP) GenerateDumps(ctx context.Context) (*dagger.Directory, error) {
	mysql := m.Build().MySQLContainer().AsService()

	_, _, err := m.bootstrapAM(ctx, mysql, UseCached) // ForceDrop not needed since it's a fresh MySQL instance.
	if err != nil {
		return nil, err
	}

	return dumpDB(mysql, mcpDBName, ssDBName)
}

// Run the e2e tests.
//
// This function configures
func (m *CCP) Etoe(
	ctx context.Context,
	// +optional
	test string,
	// +default="USE_CACHED"
	dbMode DatabaseExecutionMode,
) error {
	mysql := m.Build().MySQLContainer().
		WithMountedCache("/var/lib/mysql", dag.CacheVolume("mysql")).
		AsService()

	storage, err := m.bootstrapStorage(ctx, mysql, dbMode)
	if err != nil {
		return err
	}

	dashboard, err := m.bootstrapDashboard(ctx, mysql, storage, dbMode)
	if err != nil {
		return err
	}

	ccp := m.bootstrapCCP(mysql, storage)

	// TODO:
	// - Bootstrap CCP.
	// - Bootstrap MCPClient.

	var args []string
	{
		args = []string{"go", "test", "-v"}
		if test != "" {
			args = append(args, "-run", fmt.Sprintf("Test%s", test))
		}
		args = append(args, "./e2e/...")
	}

	dag.Go(dagger.GoOpts{
		Container: goModule().
			WithSource(m.Source.Directory("hack/ccp")).
			Container().
			WithServiceBinding("mysql", mysql).
			WithServiceBinding("dashboard", dashboard).
			WithServiceBinding("storage", storage).
			WithServiceBinding("ccp", ccp),
	}).
		Exec(args).
		Stdout(ctx)

	return nil
}

func (m *CCP) bootstrapCCP(mysql, storage *dagger.Service) *dagger.Service {
	return m.Build().CCPImage().
		WithServiceBinding("mysql", mysql).
		WithServiceBinding("storage", storage).
		AsService()
}

func (m *CCP) bootstrapAM(ctx context.Context, mysql *dagger.Service, dbMode DatabaseExecutionMode) (storage *dagger.Service, dashboard *dagger.Service, err error) {
	storage, err = m.bootstrapStorage(ctx, mysql, dbMode)
	if err != nil {
		return nil, nil, err
	}

	dashboard, err = m.bootstrapDashboard(ctx, mysql, storage, dbMode)
	if err != nil {
		return nil, nil, err
	}

	return storage, dashboard, nil
}

func (m *CCP) bootstrapStorage(ctx context.Context, mysql *dagger.Service, dbMode DatabaseExecutionMode) (*dagger.Service, error) {
	storageCtr := m.Build().StorageImage().
		WithServiceBinding("mysql", mysql).
		WithEnvVariable("DJANGO_SETTINGS_MODULE", "storage_service.settings.local").
		WithEnvVariable("SS_DB_URL", "mysql://root:12345@mysql/"+ssDBName).
		WithExposedPort(8000)

	drop := dbMode != UseCached
	if err := createDB(ctx, mysql, ssDBName, drop); err != nil {
		return nil, err
	}

	if dbMode == UseDumps {
		dumpFile := m.Source.File(filepath.Join(dumpsDir, fmt.Sprintf("%s.sql.bz2", ssDBName)))
		if err := loadDump(ctx, mysql, ssDBName, dumpFile); err != nil {
			return nil, err
		}
	}

	onlyMigrate := dbMode == UseDumps || dbMode == UseCached
	if err := bootstrapSSDB(ctx, storageCtr, onlyMigrate); err != nil {
		return nil, err
	}

	return storageCtr.AsService(), nil
}

func (m *CCP) bootstrapDashboard(ctx context.Context, mysql, storage *dagger.Service, dbMode DatabaseExecutionMode) (*dagger.Service, error) {
	dashboardCtr := m.Build().DashboardImage().
		WithServiceBinding("mysql", mysql).
		WithServiceBinding("storage", storage).
		WithEnvVariable("DJANGO_SETTINGS_MODULE", "settings.local").
		WithEnvVariable("ARCHIVEMATICA_DASHBOARD_CLIENT_USER", "root").
		WithEnvVariable("ARCHIVEMATICA_DASHBOARD_CLIENT_PASSWORD", "12345").
		WithEnvVariable("ARCHIVEMATICA_DASHBOARD_CLIENT_HOST", "mysql").
		WithEnvVariable("ARCHIVEMATICA_DASHBOARD_CLIENT_DATABASE", mcpDBName).
		WithEnvVariable("ARCHIVEMATICA_DASHBOARD_SEARCH_ENABLED", "false").
		WithExposedPort(8000)

	drop := dbMode != UseCached
	if err := createDB(ctx, mysql, mcpDBName, drop); err != nil {
		return nil, err
	}

	if dbMode == UseDumps {
		dumpFile := m.Source.File(filepath.Join(dumpsDir, fmt.Sprintf("%s.sql.bz2", mcpDBName)))
		if err := loadDump(ctx, mysql, mcpDBName, dumpFile); err != nil {
			return nil, err
		}
	}

	onlyMigrate := dbMode == UseDumps || dbMode == UseCached
	if err := bootstrapMCPDB(ctx, dashboardCtr, onlyMigrate); err != nil {
		return nil, err
	}

	return dashboardCtr.AsService(), nil
}

func createDB(ctx context.Context, mysql *dagger.Service, dbname string, drop bool) error {
	if drop {
		if err := mysqlCommand(ctx, mysql, "DROP DATABASE IF EXISTS "+dbname); err != nil {
			return err
		}
	}

	return mysqlCommand(ctx, mysql, "CREATE DATABASE IF NOT EXISTS "+dbname)
}

func mysqlCommand(ctx context.Context, mysql *dagger.Service, cmd string) error {
	_, err := dag.Container().
		From(mysqlImage).
		WithEnvVariable("CACHEBUSTER", time.Now().Format(time.RFC3339Nano)).
		WithServiceBinding("mysql", mysql).
		WithExec([]string{"mysql", "-hmysql", "-uroot", "-p12345", "-e", cmd}).
		Sync(ctx)

	return err
}

func loadDump(ctx context.Context, mysql *dagger.Service, dbname string, dump *dagger.File) error {
	_, err := dag.Container().
		From(mysqlImage).
		WithEnvVariable("CACHEBUSTER", time.Now().Format(time.RFC3339Nano)).
		WithServiceBinding("mysql", mysql).
		WithFile("/tmp/dump.sql.bz2", dump).
		WithExec([]string{"/bin/sh", "-c", "bunzip2 < /tmp/dump.sql.bz2 | mysql -hmysql -uroot -p12345 " + dbname}).
		Sync(ctx)

	return err
}

func dumpDB(mysql *dagger.Service, dbs ...string) (*dagger.Directory, error) {
	ctr := dag.Container().
		From(mysqlImage).
		WithEnvVariable("CACHEBUSTER", time.Now().Format(time.RFC3339Nano)).
		WithServiceBinding("mysql", mysql).
		WithExec([]string{"mkdir", "/tmp/dumps"}).
		WithWorkdir("/tmp/dumps")

	for _, dbname := range dbs {
		ctr = ctr.WithExec([]string{
			"/bin/sh", "-c",
			fmt.Sprintf("mysqldump -hmysql -uroot -p12345 %s | bzip2 -c > %s", dbname, fmt.Sprintf("%s.sql.bz2", dbname)),
		})
	}

	return ctr.Directory("/tmp/dumps"), nil
}

func bootstrapMCPDB(ctx context.Context, ctr *dagger.Container, onlyMigrate bool) error {
	ctr = ctr.WithEnvVariable("CACHEBUSTER", time.Now().Format(time.RFC3339Nano))

	if _, err := ctr.
		WithExec([]string{
			"/src/src/dashboard/src/manage.py",
			"migrate",
			"--noinput",
		}).
		Sync(ctx); err != nil {
		return err
	}

	if onlyMigrate {
		return nil
	}

	if _, err := ctr.
		WithExec([]string{
			"/src/src/dashboard/src/manage.py",
			"install",
			`--username="test"`,
			`--password="test"`,
			`--email="test@test.com"`,
			`--org-name="test"`,
			`--org-id="test"`,
			`--api-key="test"`,
			`--ss-url=http://storage:8000`,
			`--ss-user="test"`,
			`--ss-api-key="test"`,
			`--site-url=http://dashboard:8000`,
		}).
		Sync(ctx); err != nil {
		return err
	}

	return nil
}

func bootstrapSSDB(ctx context.Context, ctr *dagger.Container, onlyMigrate bool) error {
	ctr = ctr.WithEnvVariable("CACHEBUSTER", time.Now().Format(time.RFC3339Nano))

	if _, err := ctr.
		WithExec([]string{
			"/src/storage_service/manage.py",
			"migrate",
			"--noinput",
		}).
		Sync(ctx); err != nil {
		return err
	}

	if onlyMigrate {
		return nil
	}

	if _, err := ctr.
		WithExec([]string{
			"/src/storage_service/manage.py",
			"create_user",
			`--username="test"`,
			`--password="test"`,
			`--email="test@test.com"`,
			`--api-key="test"`,
			`--superuser`,
		}).
		Sync(ctx); err != nil {
		return err
	}

	return nil
}

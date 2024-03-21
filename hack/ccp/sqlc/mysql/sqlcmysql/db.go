// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sqlcmysql

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.cleanUpActiveJobsStmt, err = db.PrepareContext(ctx, cleanUpActiveJobs); err != nil {
		return nil, fmt.Errorf("error preparing query CleanUpActiveJobs: %w", err)
	}
	if q.cleanUpActiveSIPsStmt, err = db.PrepareContext(ctx, cleanUpActiveSIPs); err != nil {
		return nil, fmt.Errorf("error preparing query CleanUpActiveSIPs: %w", err)
	}
	if q.cleanUpActiveTasksStmt, err = db.PrepareContext(ctx, cleanUpActiveTasks); err != nil {
		return nil, fmt.Errorf("error preparing query CleanUpActiveTasks: %w", err)
	}
	if q.cleanUpActiveTransfersStmt, err = db.PrepareContext(ctx, cleanUpActiveTransfers); err != nil {
		return nil, fmt.Errorf("error preparing query CleanUpActiveTransfers: %w", err)
	}
	if q.cleanUpAwaitingJobsStmt, err = db.PrepareContext(ctx, cleanUpAwaitingJobs); err != nil {
		return nil, fmt.Errorf("error preparing query CleanUpAwaitingJobs: %w", err)
	}
	if q.cleanUpTasksWithAwaitingJobsStmt, err = db.PrepareContext(ctx, cleanUpTasksWithAwaitingJobs); err != nil {
		return nil, fmt.Errorf("error preparing query CleanUpTasksWithAwaitingJobs: %w", err)
	}
	if q.createJobStmt, err = db.PrepareContext(ctx, createJob); err != nil {
		return nil, fmt.Errorf("error preparing query CreateJob: %w", err)
	}
	if q.createWorkflowUnitVariableStmt, err = db.PrepareContext(ctx, createWorkflowUnitVariable); err != nil {
		return nil, fmt.Errorf("error preparing query CreateWorkflowUnitVariable: %w", err)
	}
	if q.getLockStmt, err = db.PrepareContext(ctx, getLock); err != nil {
		return nil, fmt.Errorf("error preparing query GetLock: %w", err)
	}
	if q.readWorkflowUnitVariableStmt, err = db.PrepareContext(ctx, readWorkflowUnitVariable); err != nil {
		return nil, fmt.Errorf("error preparing query ReadWorkflowUnitVariable: %w", err)
	}
	if q.releaseLockStmt, err = db.PrepareContext(ctx, releaseLock); err != nil {
		return nil, fmt.Errorf("error preparing query ReleaseLock: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.cleanUpActiveJobsStmt != nil {
		if cerr := q.cleanUpActiveJobsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing cleanUpActiveJobsStmt: %w", cerr)
		}
	}
	if q.cleanUpActiveSIPsStmt != nil {
		if cerr := q.cleanUpActiveSIPsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing cleanUpActiveSIPsStmt: %w", cerr)
		}
	}
	if q.cleanUpActiveTasksStmt != nil {
		if cerr := q.cleanUpActiveTasksStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing cleanUpActiveTasksStmt: %w", cerr)
		}
	}
	if q.cleanUpActiveTransfersStmt != nil {
		if cerr := q.cleanUpActiveTransfersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing cleanUpActiveTransfersStmt: %w", cerr)
		}
	}
	if q.cleanUpAwaitingJobsStmt != nil {
		if cerr := q.cleanUpAwaitingJobsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing cleanUpAwaitingJobsStmt: %w", cerr)
		}
	}
	if q.cleanUpTasksWithAwaitingJobsStmt != nil {
		if cerr := q.cleanUpTasksWithAwaitingJobsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing cleanUpTasksWithAwaitingJobsStmt: %w", cerr)
		}
	}
	if q.createJobStmt != nil {
		if cerr := q.createJobStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createJobStmt: %w", cerr)
		}
	}
	if q.createWorkflowUnitVariableStmt != nil {
		if cerr := q.createWorkflowUnitVariableStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createWorkflowUnitVariableStmt: %w", cerr)
		}
	}
	if q.getLockStmt != nil {
		if cerr := q.getLockStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLockStmt: %w", cerr)
		}
	}
	if q.readWorkflowUnitVariableStmt != nil {
		if cerr := q.readWorkflowUnitVariableStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing readWorkflowUnitVariableStmt: %w", cerr)
		}
	}
	if q.releaseLockStmt != nil {
		if cerr := q.releaseLockStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing releaseLockStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                               DBTX
	tx                               *sql.Tx
	cleanUpActiveJobsStmt            *sql.Stmt
	cleanUpActiveSIPsStmt            *sql.Stmt
	cleanUpActiveTasksStmt           *sql.Stmt
	cleanUpActiveTransfersStmt       *sql.Stmt
	cleanUpAwaitingJobsStmt          *sql.Stmt
	cleanUpTasksWithAwaitingJobsStmt *sql.Stmt
	createJobStmt                    *sql.Stmt
	createWorkflowUnitVariableStmt   *sql.Stmt
	getLockStmt                      *sql.Stmt
	readWorkflowUnitVariableStmt     *sql.Stmt
	releaseLockStmt                  *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                               tx,
		tx:                               tx,
		cleanUpActiveJobsStmt:            q.cleanUpActiveJobsStmt,
		cleanUpActiveSIPsStmt:            q.cleanUpActiveSIPsStmt,
		cleanUpActiveTasksStmt:           q.cleanUpActiveTasksStmt,
		cleanUpActiveTransfersStmt:       q.cleanUpActiveTransfersStmt,
		cleanUpAwaitingJobsStmt:          q.cleanUpAwaitingJobsStmt,
		cleanUpTasksWithAwaitingJobsStmt: q.cleanUpTasksWithAwaitingJobsStmt,
		createJobStmt:                    q.createJobStmt,
		createWorkflowUnitVariableStmt:   q.createWorkflowUnitVariableStmt,
		getLockStmt:                      q.getLockStmt,
		readWorkflowUnitVariableStmt:     q.readWorkflowUnitVariableStmt,
		releaseLockStmt:                  q.releaseLockStmt,
	}
}

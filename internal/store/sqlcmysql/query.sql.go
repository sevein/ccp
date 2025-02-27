// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package sqlcmysql

import (
	"context"
	"database/sql"
	"time"

	uuid "github.com/google/uuid"
)

const cleanUpActiveJobs = `-- name: CleanUpActiveJobs :exec
UPDATE Jobs SET currentStep = 4 WHERE currentStep = 3
`

func (q *Queries) CleanUpActiveJobs(ctx context.Context) error {
	_, err := q.exec(ctx, q.cleanUpActiveJobsStmt, cleanUpActiveJobs)
	return err
}

const cleanUpActiveSIPs = `-- name: CleanUpActiveSIPs :exec
UPDATE SIPs SET status = 4, completed_at = UTC_TIMESTAMP() WHERE status IN (0, 1)
`

func (q *Queries) CleanUpActiveSIPs(ctx context.Context) error {
	_, err := q.exec(ctx, q.cleanUpActiveSIPsStmt, cleanUpActiveSIPs)
	return err
}

const cleanUpActiveTasks = `-- name: CleanUpActiveTasks :exec
UPDATE Tasks SET exitCode = -1, stdError = "MCP shut down while processing." WHERE exitCode IS NULL
`

func (q *Queries) CleanUpActiveTasks(ctx context.Context) error {
	_, err := q.exec(ctx, q.cleanUpActiveTasksStmt, cleanUpActiveTasks)
	return err
}

const cleanUpActiveTransfers = `-- name: CleanUpActiveTransfers :exec
UPDATE Transfers SET status = 4, completed_at = UTC_TIMESTAMP() WHERE status IN (0, 1)
`

func (q *Queries) CleanUpActiveTransfers(ctx context.Context) error {
	_, err := q.exec(ctx, q.cleanUpActiveTransfersStmt, cleanUpActiveTransfers)
	return err
}

const cleanUpAwaitingJobs = `-- name: CleanUpAwaitingJobs :exec
DELETE FROM Jobs WHERE currentStep = 1
`

func (q *Queries) CleanUpAwaitingJobs(ctx context.Context) error {
	_, err := q.exec(ctx, q.cleanUpAwaitingJobsStmt, cleanUpAwaitingJobs)
	return err
}

const cleanUpTasksWithAwaitingJobs = `-- name: CleanUpTasksWithAwaitingJobs :exec

DELETE FROM Tasks WHERE jobuuid IN (SELECT jobUUID FROM Jobs WHERE currentStep = 1)
`

// Clean-ups
func (q *Queries) CleanUpTasksWithAwaitingJobs(ctx context.Context) error {
	_, err := q.exec(ctx, q.cleanUpTasksWithAwaitingJobsStmt, cleanUpTasksWithAwaitingJobs)
	return err
}

const createJob = `-- name: CreateJob :exec

INSERT INTO Jobs (jobUUID, jobType, createdTime, createdTimeDec, directory, SIPUUID, unitType, currentStep, microserviceGroup, hidden, MicroServiceChainLinksPK, subJobOf) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateJobParams struct {
	ID                uuid.UUID
	Type              string
	CreatedAt         time.Time
	Createdtimedec    string
	Directory         string
	SIPID             uuid.UUID
	Unittype          string
	Currentstep       int32
	Microservicegroup string
	Hidden            bool
	LinkID            uuid.NullUUID
	Subjobof          string
}

// Jobs
func (q *Queries) CreateJob(ctx context.Context, arg *CreateJobParams) error {
	_, err := q.exec(ctx, q.createJobStmt, createJob,
		arg.ID,
		arg.Type,
		arg.CreatedAt,
		arg.Createdtimedec,
		arg.Directory,
		arg.SIPID,
		arg.Unittype,
		arg.Currentstep,
		arg.Microservicegroup,
		arg.Hidden,
		arg.LinkID,
		arg.Subjobof,
	)
	return err
}

const createSIP = `-- name: CreateSIP :exec

INSERT INTO SIPs (sipUUID, createdTime, currentPath, hidden, aipFilename, sipType, dirUUIDs, status, completed_at) VALUES (?, UTC_TIMESTAMP(), ?, 0, '', ?, 0, 0, NULL)
`

type CreateSIPParams struct {
	SIPID       uuid.UUID
	Currentpath sql.NullString
	Siptype     string
}

// SIPs
func (q *Queries) CreateSIP(ctx context.Context, arg *CreateSIPParams) error {
	_, err := q.exec(ctx, q.createSIPStmt, createSIP, arg.SIPID, arg.Currentpath, arg.Siptype)
	return err
}

const createTransfer = `-- name: CreateTransfer :exec

INSERT INTO Transfers (transferUUID, currentLocation, type, accessionID, sourceOfAcquisition, typeOfTransfer, description, notes, access_system_id, hidden, transferMetadataSetRowUUID, dirUUIDs, status, completed_at)
VALUES (?, ?, '', ?, '', '', '', '', ?, 0, ?, 0, 0, NULL)
`

type CreateTransferParams struct {
	Transferuuid               uuid.UUID
	Currentlocation            string
	Accessionid                string
	AccessSystemID             string
	Transfermetadatasetrowuuid uuid.NullUUID
}

// Transfers
func (q *Queries) CreateTransfer(ctx context.Context, arg *CreateTransferParams) error {
	_, err := q.exec(ctx, q.createTransferStmt, createTransfer,
		arg.Transferuuid,
		arg.Currentlocation,
		arg.Accessionid,
		arg.AccessSystemID,
		arg.Transfermetadatasetrowuuid,
	)
	return err
}

const createUnitVar = `-- name: CreateUnitVar :exec
INSERT INTO UnitVariables (pk, unitType, unitUUID, variable, variableValue, microServiceChainLink, createdTime, updatedTime)
VALUES (
    UUID(),
    ?,
    ?,
    ?,
    ?,
    ?,
    UTC_TIMESTAMP(),
    UTC_TIMESTAMP()
)
`

type CreateUnitVarParams struct {
	UnitType sql.NullString
	UnitID   uuid.UUID
	Name     sql.NullString
	Value    sql.NullString
	LinkID   uuid.NullUUID
}

func (q *Queries) CreateUnitVar(ctx context.Context, arg *CreateUnitVarParams) error {
	_, err := q.exec(ctx, q.createUnitVarStmt, createUnitVar,
		arg.UnitType,
		arg.UnitID,
		arg.Name,
		arg.Value,
		arg.LinkID,
	)
	return err
}

const listJobs = `-- name: ListJobs :many
SELECT jobuuid, jobtype, createdtime, createdtimedec, directory, sipuuid, unittype, currentstep, microservicegroup, hidden, subjobof, microservicechainlinkspk FROM Jobs WHERE SIPUUID = ? ORDER BY createdTime DESC
`

func (q *Queries) ListJobs(ctx context.Context, sipuuid uuid.UUID) ([]*Job, error) {
	rows, err := q.query(ctx, q.listJobsStmt, listJobs, sipuuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Job{}
	for rows.Next() {
		var i Job
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.CreatedAt,
			&i.Createdtimedec,
			&i.Directory,
			&i.SIPID,
			&i.Unittype,
			&i.Currentstep,
			&i.Microservicegroup,
			&i.Hidden,
			&i.Subjobof,
			&i.LinkID,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSIPsWithCreationTimestamps = `-- name: ListSIPsWithCreationTimestamps :many
SELECT
    j.SIPUUID,
    j.createdTime AS created_at,
    j.createdTimeDec AS created_at_dec,
    s.status
FROM Jobs j
JOIN (
    SELECT
        SIPUUID,
        MAX(createdTime) AS max_created_at
    FROM Jobs
    WHERE unitType = 'unitSIP' AND NOT SIPUUID LIKE '%None%'
    GROUP BY SIPUUID
) AS latest_jobs ON j.SIPUUID = latest_jobs.SIPUUID AND j.createdTime = latest_jobs.max_created_at
LEFT JOIN SIPs s ON s.sipUUID = j.SIPUUID
WHERE j.unitType = 'unitSIP' AND NOT j.SIPUUID LIKE '%None%' AND s.hidden = 0
`

type ListSIPsWithCreationTimestampsRow struct {
	SIPID        uuid.UUID
	CreatedAt    time.Time
	CreatedAtDec string
	Status       sql.NullInt16
}

func (q *Queries) ListSIPsWithCreationTimestamps(ctx context.Context) ([]*ListSIPsWithCreationTimestampsRow, error) {
	rows, err := q.query(ctx, q.listSIPsWithCreationTimestampsStmt, listSIPsWithCreationTimestamps)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*ListSIPsWithCreationTimestampsRow{}
	for rows.Next() {
		var i ListSIPsWithCreationTimestampsRow
		if err := rows.Scan(
			&i.SIPID,
			&i.CreatedAt,
			&i.CreatedAtDec,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTransfersWithCreationTimestamps = `-- name: ListTransfersWithCreationTimestamps :many
SELECT
    j.SIPUUID,
    j.createdTime AS created_at,
    j.createdTimeDec AS created_at_dec,
    t.status
FROM Jobs j
JOIN (
    SELECT
        SIPUUID,
        MAX(createdTime) AS max_created_at
    FROM Jobs
    WHERE unitType = 'unitTransfer' AND NOT SIPUUID LIKE '%None%'
    GROUP BY SIPUUID
) AS latest_jobs ON j.SIPUUID = latest_jobs.SIPUUID AND j.createdTime = latest_jobs.max_created_at
LEFT JOIN Transfers t ON t.transferUUID = j.SIPUUID
WHERE j.unitType = 'unitTransfer' AND NOT j.SIPUUID LIKE '%None%' AND t.hidden = 0
`

type ListTransfersWithCreationTimestampsRow struct {
	SIPID        uuid.UUID
	CreatedAt    time.Time
	CreatedAtDec string
	Status       sql.NullInt16
}

func (q *Queries) ListTransfersWithCreationTimestamps(ctx context.Context) ([]*ListTransfersWithCreationTimestampsRow, error) {
	rows, err := q.query(ctx, q.listTransfersWithCreationTimestampsStmt, listTransfersWithCreationTimestamps)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*ListTransfersWithCreationTimestampsRow{}
	for rows.Next() {
		var i ListTransfersWithCreationTimestampsRow
		if err := rows.Scan(
			&i.SIPID,
			&i.CreatedAt,
			&i.CreatedAtDec,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const readDashboardSetting = `-- name: ReadDashboardSetting :one
SELECT name, value, scope FROM DashboardSettings WHERE name = ?
`

type ReadDashboardSettingRow struct {
	Name  string
	Value string
	Scope string
}

func (q *Queries) ReadDashboardSetting(ctx context.Context, name string) (*ReadDashboardSettingRow, error) {
	row := q.queryRow(ctx, q.readDashboardSettingStmt, readDashboardSetting, name)
	var i ReadDashboardSettingRow
	err := row.Scan(&i.Name, &i.Value, &i.Scope)
	return &i, err
}

const readDashboardSettingsWithNameLike = `-- name: ReadDashboardSettingsWithNameLike :many
SELECT name, value, scope FROM DashboardSettings WHERE name LIKE ?
`

type ReadDashboardSettingsWithNameLikeRow struct {
	Name  string
	Value string
	Scope string
}

func (q *Queries) ReadDashboardSettingsWithNameLike(ctx context.Context, name string) ([]*ReadDashboardSettingsWithNameLikeRow, error) {
	rows, err := q.query(ctx, q.readDashboardSettingsWithNameLikeStmt, readDashboardSettingsWithNameLike, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*ReadDashboardSettingsWithNameLikeRow{}
	for rows.Next() {
		var i ReadDashboardSettingsWithNameLikeRow
		if err := rows.Scan(&i.Name, &i.Value, &i.Scope); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const readDashboardSettingsWithScope = `-- name: ReadDashboardSettingsWithScope :many

SELECT name, value, scope FROM DashboardSettings WHERE scope = ?
`

type ReadDashboardSettingsWithScopeRow struct {
	Name  string
	Value string
	Scope string
}

// Dashboard settings
func (q *Queries) ReadDashboardSettingsWithScope(ctx context.Context, scope string) ([]*ReadDashboardSettingsWithScopeRow, error) {
	rows, err := q.query(ctx, q.readDashboardSettingsWithScopeStmt, readDashboardSettingsWithScope, scope)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*ReadDashboardSettingsWithScopeRow{}
	for rows.Next() {
		var i ReadDashboardSettingsWithScopeRow
		if err := rows.Scan(&i.Name, &i.Value, &i.Scope); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const readSIP = `-- name: ReadSIP :one
SELECT sipUUID, createdTime, currentPath, hidden, aipFilename, sipType, dirUUIDs, status, completed_at FROM SIPs WHERE sipUUID = ?
`

type ReadSIPRow struct {
	SIPID       uuid.UUID
	CreatedAt   time.Time
	Currentpath sql.NullString
	Hidden      bool
	Aipfilename sql.NullString
	Siptype     string
	Diruuids    bool
	Status      uint16
	CompletedAt sql.NullTime
}

func (q *Queries) ReadSIP(ctx context.Context, sipuuid uuid.UUID) (*ReadSIPRow, error) {
	row := q.queryRow(ctx, q.readSIPStmt, readSIP, sipuuid)
	var i ReadSIPRow
	err := row.Scan(
		&i.SIPID,
		&i.CreatedAt,
		&i.Currentpath,
		&i.Hidden,
		&i.Aipfilename,
		&i.Siptype,
		&i.Diruuids,
		&i.Status,
		&i.CompletedAt,
	)
	return &i, err
}

const readSIPLocation = `-- name: ReadSIPLocation :one
SELECT sipUUID, currentPath FROM SIPs WHERE sipUUID = ?
`

type ReadSIPLocationRow struct {
	SIPID       uuid.UUID
	Currentpath sql.NullString
}

func (q *Queries) ReadSIPLocation(ctx context.Context, sipuuid uuid.UUID) (*ReadSIPLocationRow, error) {
	row := q.queryRow(ctx, q.readSIPLocationStmt, readSIPLocation, sipuuid)
	var i ReadSIPLocationRow
	err := row.Scan(&i.SIPID, &i.Currentpath)
	return &i, err
}

const readSIPWithLocation = `-- name: ReadSIPWithLocation :one
SELECT sipUUID FROM SIPs WHERE currentPath = ?
`

func (q *Queries) ReadSIPWithLocation(ctx context.Context, currentpath sql.NullString) (uuid.UUID, error) {
	row := q.queryRow(ctx, q.readSIPWithLocationStmt, readSIPWithLocation, currentpath)
	var sipuuid uuid.UUID
	err := row.Scan(&sipuuid)
	return sipuuid, err
}

const readTransfer = `-- name: ReadTransfer :one
SELECT transferUUID, currentLocation, type, accessionID, sourceOfAcquisition, typeOfTransfer, description, notes, access_system_id, hidden, transferMetadataSetRowUUID, dirUUIDs, status, completed_at FROM Transfers WHERE transferUUID = ?
`

type ReadTransferRow struct {
	Transferuuid               uuid.UUID
	Currentlocation            string
	Type                       string
	Accessionid                string
	Sourceofacquisition        string
	Typeoftransfer             string
	Description                string
	Notes                      string
	AccessSystemID             string
	Hidden                     bool
	Transfermetadatasetrowuuid uuid.NullUUID
	Diruuids                   bool
	Status                     uint16
	CompletedAt                sql.NullTime
}

func (q *Queries) ReadTransfer(ctx context.Context, transferuuid uuid.UUID) (*ReadTransferRow, error) {
	row := q.queryRow(ctx, q.readTransferStmt, readTransfer, transferuuid)
	var i ReadTransferRow
	err := row.Scan(
		&i.Transferuuid,
		&i.Currentlocation,
		&i.Type,
		&i.Accessionid,
		&i.Sourceofacquisition,
		&i.Typeoftransfer,
		&i.Description,
		&i.Notes,
		&i.AccessSystemID,
		&i.Hidden,
		&i.Transfermetadatasetrowuuid,
		&i.Diruuids,
		&i.Status,
		&i.CompletedAt,
	)
	return &i, err
}

const readTransferLocation = `-- name: ReadTransferLocation :one
SELECT transferUUID, currentLocation FROM Transfers WHERE transferUUID = ?
`

type ReadTransferLocationRow struct {
	Transferuuid    uuid.UUID
	Currentlocation string
}

func (q *Queries) ReadTransferLocation(ctx context.Context, transferuuid uuid.UUID) (*ReadTransferLocationRow, error) {
	row := q.queryRow(ctx, q.readTransferLocationStmt, readTransferLocation, transferuuid)
	var i ReadTransferLocationRow
	err := row.Scan(&i.Transferuuid, &i.Currentlocation)
	return &i, err
}

const readTransferWithLocation = `-- name: ReadTransferWithLocation :one
SELECT transferUUID FROM Transfers WHERE currentLocation = ?
`

func (q *Queries) ReadTransferWithLocation(ctx context.Context, currentlocation string) (uuid.UUID, error) {
	row := q.queryRow(ctx, q.readTransferWithLocationStmt, readTransferWithLocation, currentlocation)
	var transferuuid uuid.UUID
	err := row.Scan(&transferuuid)
	return transferuuid, err
}

const readUnitVar = `-- name: ReadUnitVar :one

SELECT variableValue, microServiceChainLink FROM UnitVariables WHERE unitType = ? AND unitUUID = ? AND variable = ?
`

type ReadUnitVarParams struct {
	UnitType sql.NullString
	UnitID   uuid.UUID
	Name     sql.NullString
}

type ReadUnitVarRow struct {
	Variablevalue sql.NullString
	LinkID        uuid.NullUUID
}

// Unit variables
func (q *Queries) ReadUnitVar(ctx context.Context, arg *ReadUnitVarParams) (*ReadUnitVarRow, error) {
	row := q.queryRow(ctx, q.readUnitVarStmt, readUnitVar, arg.UnitType, arg.UnitID, arg.Name)
	var i ReadUnitVarRow
	err := row.Scan(&i.Variablevalue, &i.LinkID)
	return &i, err
}

const readUnitVars = `-- name: ReadUnitVars :many
SELECT unitType, unitUUID, variable, variableValue, microServiceChainLink FROM UnitVariables WHERE unitUUID = ? AND variable = ?
`

type ReadUnitVarsParams struct {
	UnitID uuid.UUID
	Name   sql.NullString
}

type ReadUnitVarsRow struct {
	Unittype      sql.NullString
	Unituuid      uuid.UUID
	Variable      sql.NullString
	Variablevalue sql.NullString
	LinkID        uuid.NullUUID
}

func (q *Queries) ReadUnitVars(ctx context.Context, arg *ReadUnitVarsParams) ([]*ReadUnitVarsRow, error) {
	rows, err := q.query(ctx, q.readUnitVarsStmt, readUnitVars, arg.UnitID, arg.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*ReadUnitVarsRow{}
	for rows.Next() {
		var i ReadUnitVarsRow
		if err := rows.Scan(
			&i.Unittype,
			&i.Unituuid,
			&i.Variable,
			&i.Variablevalue,
			&i.LinkID,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const readUserWithKey = `-- name: ReadUserWithKey :one

SELECT auth_user.id, auth_user.username, auth_user.email, auth_user.is_active, main_userprofile.agent_id
FROM auth_user
JOIN tastypie_apikey ON auth_user.id = tastypie_apikey.user_id
LEFT JOIN main_userprofile ON auth_user.id = main_userprofile.user_id
WHERE auth_user.username = ? AND tastypie_apikey.key = ? AND auth_user.is_active = 1
LIMIT 1
`

type ReadUserWithKeyParams struct {
	Username string
	Key      string
}

type ReadUserWithKeyRow struct {
	ID       int32
	Username string
	Email    string
	IsActive bool
	AgentID  sql.NullInt32
}

// Authorization
func (q *Queries) ReadUserWithKey(ctx context.Context, arg *ReadUserWithKeyParams) (*ReadUserWithKeyRow, error) {
	row := q.queryRow(ctx, q.readUserWithKeyStmt, readUserWithKey, arg.Username, arg.Key)
	var i ReadUserWithKeyRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.IsActive,
		&i.AgentID,
	)
	return &i, err
}

const updateJobStatus = `-- name: UpdateJobStatus :exec
UPDATE Jobs SET currentStep = ? WHERE jobUUID = ?
`

type UpdateJobStatusParams struct {
	Currentstep int32
	ID          uuid.UUID
}

func (q *Queries) UpdateJobStatus(ctx context.Context, arg *UpdateJobStatusParams) error {
	_, err := q.exec(ctx, q.updateJobStatusStmt, updateJobStatus, arg.Currentstep, arg.ID)
	return err
}

const updateSIPLocation = `-- name: UpdateSIPLocation :exec
UPDATE SIPs SET currentPath = ? WHERE sipUUID = ?
`

type UpdateSIPLocationParams struct {
	Currentpath sql.NullString
	SIPID       uuid.UUID
}

func (q *Queries) UpdateSIPLocation(ctx context.Context, arg *UpdateSIPLocationParams) error {
	_, err := q.exec(ctx, q.updateSIPLocationStmt, updateSIPLocation, arg.Currentpath, arg.SIPID)
	return err
}

const updateSIPStatus = `-- name: UpdateSIPStatus :exec
UPDATE SIPs SET status = ? WHERE sipUUID = ?
`

type UpdateSIPStatusParams struct {
	Status uint16
	SIPID  uuid.UUID
}

func (q *Queries) UpdateSIPStatus(ctx context.Context, arg *UpdateSIPStatusParams) error {
	_, err := q.exec(ctx, q.updateSIPStatusStmt, updateSIPStatus, arg.Status, arg.SIPID)
	return err
}

const updateTransferLocation = `-- name: UpdateTransferLocation :exec
UPDATE Transfers SET currentLocation = ? WHERE transferUUID = ?
`

type UpdateTransferLocationParams struct {
	Currentlocation string
	Transferuuid    uuid.UUID
}

func (q *Queries) UpdateTransferLocation(ctx context.Context, arg *UpdateTransferLocationParams) error {
	_, err := q.exec(ctx, q.updateTransferLocationStmt, updateTransferLocation, arg.Currentlocation, arg.Transferuuid)
	return err
}

const updateTransferStatus = `-- name: UpdateTransferStatus :exec
UPDATE Transfers SET status = ? WHERE transferUUID = ?
`

type UpdateTransferStatusParams struct {
	Status       uint16
	Transferuuid uuid.UUID
}

func (q *Queries) UpdateTransferStatus(ctx context.Context, arg *UpdateTransferStatusParams) error {
	_, err := q.exec(ctx, q.updateTransferStatusStmt, updateTransferStatus, arg.Status, arg.Transferuuid)
	return err
}

const updateUnitVar = `-- name: UpdateUnitVar :exec
UPDATE UnitVariables
SET
    variableValue = ?,
    microServiceChainLink = ?,
    updatedTime = UTC_TIMESTAMP()
WHERE
    unitType = ?
    AND unitUUID = ?
    AND variable = ?
`

type UpdateUnitVarParams struct {
	Value    sql.NullString
	LinkID   uuid.NullUUID
	UnitType sql.NullString
	UnitID   uuid.UUID
	Name     sql.NullString
}

func (q *Queries) UpdateUnitVar(ctx context.Context, arg *UpdateUnitVarParams) error {
	_, err := q.exec(ctx, q.updateUnitVarStmt, updateUnitVar,
		arg.Value,
		arg.LinkID,
		arg.UnitType,
		arg.UnitID,
		arg.Name,
	)
	return err
}

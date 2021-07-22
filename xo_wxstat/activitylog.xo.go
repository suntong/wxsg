package xo_wxstat

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// ActivityLog represents a row from 'activity_log'.
type ActivityLog struct {
	ActivityLogID  int  `json:"activity_log_id"`  // activity_log_id
	UserID         int  `json:"user_id"`          // user_id
	GroupID        int  `json:"group_id"`         // group_id
	ActivityTypeID int  `json:"activity_type_id"` // activity_type_id
	ActivityDate   Time `json:"activity_date"`    // activity_date
	Stat           int  `json:"stat"`             // stat
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the ActivityLog exists in the database.
func (al *ActivityLog) Exists() bool {
	return al._exists
}

// Deleted returns true when the ActivityLog has been marked for deletion from
// the database.
func (al *ActivityLog) Deleted() bool {
	return al._deleted
}

// Insert inserts the ActivityLog to the database.
func (al *ActivityLog) Insert(ctx context.Context, db DB) error {
	switch {
	case al._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case al._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO activity_log (` +
		`activity_log_id, user_id, group_id, activity_type_id, activity_date, stat` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`)`
	// run
	logf(sqlstr, al.ActivityLogID, al.UserID, al.GroupID, al.ActivityTypeID, al.ActivityDate, al.Stat)
	if _, err := db.ExecContext(ctx, sqlstr, al.ActivityLogID, al.UserID, al.GroupID, al.ActivityTypeID, al.ActivityDate, al.Stat); err != nil {
		return logerror(err)
	}
	// set exists
	al._exists = true
	return nil
}

// Update updates a ActivityLog in the database.
func (al *ActivityLog) Update(ctx context.Context, db DB) error {
	switch {
	case !al._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case al._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE activity_log SET ` +
		`user_id = $1, group_id = $2, activity_type_id = $3, activity_date = $4, stat = $5 ` +
		`WHERE activity_log_id = $6`
	// run
	logf(sqlstr, al.UserID, al.GroupID, al.ActivityTypeID, al.ActivityDate, al.Stat, al.ActivityLogID)
	if _, err := db.ExecContext(ctx, sqlstr, al.UserID, al.GroupID, al.ActivityTypeID, al.ActivityDate, al.Stat, al.ActivityLogID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the ActivityLog to the database.
func (al *ActivityLog) Save(ctx context.Context, db DB) error {
	if al.Exists() {
		return al.Update(ctx, db)
	}
	return al.Insert(ctx, db)
}

// Upsert performs an upsert for ActivityLog.
func (al *ActivityLog) Upsert(ctx context.Context, db DB) error {
	switch {
	case al._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO activity_log (` +
		`activity_log_id, user_id, group_id, activity_type_id, activity_date, stat` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`)` +
		` ON CONFLICT (activity_log_id) DO ` +
		`UPDATE SET ` +
		`user_id = EXCLUDED.user_id, group_id = EXCLUDED.group_id, activity_type_id = EXCLUDED.activity_type_id, activity_date = EXCLUDED.activity_date, stat = EXCLUDED.stat `
	// run
	logf(sqlstr, al.ActivityLogID, al.UserID, al.GroupID, al.ActivityTypeID, al.ActivityDate, al.Stat)
	if _, err := db.ExecContext(ctx, sqlstr, al.ActivityLogID, al.UserID, al.GroupID, al.ActivityTypeID, al.ActivityDate, al.Stat); err != nil {
		return err
	}
	// set exists
	al._exists = true
	return nil
}

// Delete deletes the ActivityLog from the database.
func (al *ActivityLog) Delete(ctx context.Context, db DB) error {
	switch {
	case !al._exists: // doesn't exist
		return nil
	case al._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM activity_log ` +
		`WHERE activity_log_id = $1`
	// run
	logf(sqlstr, al.ActivityLogID)
	if _, err := db.ExecContext(ctx, sqlstr, al.ActivityLogID); err != nil {
		return logerror(err)
	}
	// set deleted
	al._deleted = true
	return nil
}

// ActivityLogByActivityLogID retrieves a row from 'activity_log' as a ActivityLog.
//
// Generated from index 'activity_log_activity_log_id_pkey'.
func ActivityLogByActivityLogID(ctx context.Context, db DB, activityLogID int) (*ActivityLog, error) {
	// query
	const sqlstr = `SELECT ` +
		`activity_log_id, user_id, group_id, activity_type_id, activity_date, stat ` +
		`FROM activity_log ` +
		`WHERE activity_log_id = $1`
	// run
	logf(sqlstr, activityLogID)
	al := ActivityLog{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, activityLogID).Scan(&al.ActivityLogID, &al.UserID, &al.GroupID, &al.ActivityTypeID, &al.ActivityDate, &al.Stat); err != nil {
		return nil, logerror(err)
	}
	return &al, nil
}

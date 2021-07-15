package xowxstat

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// Custom defined search query GetGroupID
type GetGroupID struct {
	GroupID int `json:"group_id"` // group_id
}

// GetGroupIDsByGroupName runs a custom query, returning results as GetGroupID.
func GetGroupIDsByGroupName(ctx context.Context, db DB, GroupName string) ([]*GetGroupID, error) {
	// query
	const sqlstr = `SELECT group_id ` +
		`FROM wx_group ` +
		`WHERE name == '$1';`
	// run
	logf(sqlstr, GroupName)
	rows, err := db.QueryContext(ctx, sqlstr, GroupName)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*GetGroupID
	for rows.Next() {
		var gg GetGroupID
		// scan
		if err := rows.Scan(&gg.GroupID); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &gg)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

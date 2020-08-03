// Package xoWxstat contains the types for schema ''.
package xoWxstat

// Code generated by xo. DO NOT EDIT.

// Custom defined search query GetGroupID
type GetGroupID struct {
	GroupID int // group_id
}

// GetGroupIDsByGroupName runs a custom query, returning results as GetGroupID.
func GetGroupIDsByGroupName(db XODB, GroupName string) ([]*GetGroupID, error) {
	var err error

	// sql query
	const sqlstr = `SELECT group_id ` +
		`FROM wx_group ` +
		`WHERE name == '?';`

	// run query
	XOLog(sqlstr, GroupName)
	q, err := db.Query(sqlstr, GroupName)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*GetGroupID{}
	for q.Next() {
		gg := GetGroupID{}

		// scan
		err = q.Scan(&gg.GroupID)
		if err != nil {
			return nil, err
		}

		res = append(res, &gg)
	}

	return res, nil
}

// Package xoWxstat contains the types for schema ''.
package xoWxstat

// Code generated by xo. DO NOT EDIT.

// Custom defined search query GetGroupSettings
type GetGroupSettings struct {
	ActivePeriod int // active_period
	QuietPeriod  int // quiet_period
	ActiveNum    int // active_num
	QuietNum     int // quiet_num
}

// GetGroupSettingsByGroupID runs a custom query, returning results as GetGroupSettings.
func GetGroupSettingsByGroupID(db XODB, GroupID string) ([]*GetGroupSettings, error) {
	var err error

	// sql query
	const sqlstr = `SELECT active_period, quiet_period, active_num, quiet_num ` +
		`FROM wx_group ` +
		`WHERE group_id == ?;`

	// run query
	XOLog(sqlstr, GroupID)
	q, err := db.Query(sqlstr, GroupID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*GetGroupSettings{}
	for q.Next() {
		ggs := GetGroupSettings{}

		// scan
		err = q.Scan(&ggs.ActivePeriod, &ggs.QuietPeriod, &ggs.ActiveNum, &ggs.QuietNum)
		if err != nil {
			return nil, err
		}

		res = append(res, &ggs)
	}

	return res, nil
}

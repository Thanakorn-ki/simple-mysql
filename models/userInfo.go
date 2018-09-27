package models

type userInfo struct {
	uid        int
	username   string
	departname string
	created    string
}

func (db *DB) AllUserInfo() ([]*userInfo, error) {

	rows, err := db.Query("SELECT * FROM Userinfo")
	if err != nil {
		return nil, err
	}
	usersInfo := make([]*userInfo, 0)
	for rows.Next() {
		bk := new(userInfo)
		err := rows.Scan(&bk.uid, &bk.username, &bk.departname, &bk.created)
		if err != nil {
			return nil, err
		}
		usersInfo = append(usersInfo, bk)
	}
	// usersInfo := rows
	defer rows.Close()
	return usersInfo, nil
}

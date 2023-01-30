package database

func (db *appdbimpl) AddPhoto(p Photo) (Photo, error) {

	/*	var i int
		db.c.QueryRow("SELECT id FROM users WHERE token=?", tk).Scan(&i)
	*/
	res, err := db.c.Exec("INSERT INTO photos(user_id,data,photo) VALUES(?,?,?) ", p.User_id, p.Data, p.Photo)
	if err != nil {
		return Photo{}, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return p, err
	}
	p.Id_photo = uint64(lastInsertID)

	return p, nil
}

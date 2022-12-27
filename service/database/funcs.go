package database

func (db *appdbimpl) GetIdbyUsername(username string) int {
	row, err := db.c.Query("SELECT id FROM users WHERE username=?", username)
	if err != nil {
		return 0
	}
	var id int
	err = row.Scan(&id)
	if err != nil {
		return 0
	}
	return id
}

func (db *appdbimpl) GetPhotosbyId(id int) ([]Photo, error) {
	var Photos []Photo
	rows, err := db.c.Query("SELECT id_photo,user_id,data,photo FROM photos WHERE user_id=?", id)
	if err != nil {
		return Photos, err
	}
	for rows.Next() {
		p := Photo{}
		err = rows.Scan(&p.Id_photo, &p.User_id, &p.Data, &p.Photo)
		if err != nil {
			return Photos, err
		}
		Photos = append(Photos, p)
	}
	if err = rows.Err(); err != nil {
		return Photos, err
	}
	rows.Close()

	return Photos, err
}

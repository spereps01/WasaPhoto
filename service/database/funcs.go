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
func (db *appdbimpl) GetUsernamebyId(id int) string {
	var username string
	err := db.c.QueryRow("SELECT username FROM users WHERE id=?", id).Scan(&username)
	if err != nil {
		return ""
	}

	return username
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

		var count int
		err = db.c.QueryRow("SELECT COUNT(*) FROM likes WHERE id_photo=? ", p.Id_photo).Scan(&count)
		if err != nil {
			return Photos, err
		}
		p.N_like = uint64(count)

		p.Username = db.GetUsernamebyId(int(p.User_id))
		Photos = append(Photos, p)
	}
	if err = rows.Err(); err != nil {
		return Photos, err
	}
	rows.Close()

	return Photos, err
}

func (db *appdbimpl) GetFollowersbyId(id int) (int, error) {

	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM follow WHERE id2=?", id).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, err
}

func (db *appdbimpl) GetFollowingsbyId(id int) (int, error) {

	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM follow WHERE id1=?", id).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, err
}

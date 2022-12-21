package database

func (db *appdbimpl) GetUserProfile(username string) ([]Profile, error) {
	var Profiles []Profile
	rows, err := db.c.Query("SELECT username,id FROM users WHERE username LIKE ?", username)
	if err != nil {
		return Profiles, err
	}
	for rows.Next() {
		p := Profile{}
		err = rows.Scan(&p.Username, &p.Id)
		var Photos []Photo
		Photos, err = db.GetPhotosbyId(int(p.Id))
		p.Photos = Photos
		Profiles = append(Profiles, p)
	}

	return Profiles, err
}

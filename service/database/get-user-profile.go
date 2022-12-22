package database

func (db *appdbimpl) GetUserProfile(username string) ([]Profile, error) {
	var Profiles []Profile
	rows, err := db.c.Query("SELECT id,username FROM users WHERE username LIKE ?", "%"+username+"%")

	if err != nil {
		return Profiles, err
	}

	for rows.Next() {
		p := Profile{}

		err = rows.Scan(&p.Id, &p.Username)
		if err != nil {
			return Profiles, err
		}

		var Photos []Photo

		Photos, err = db.GetPhotosbyId(int(p.Id))
		if err != nil {
			return Profiles, err
		}

		p.Photos = Photos
		Profiles = append(Profiles, p)

	}

	return Profiles, err
}

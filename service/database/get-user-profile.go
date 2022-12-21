package database

import "log"

func (db *appdbimpl) GetUserProfile(username string) ([]Profile, error) {
	var Profiles []Profile
	rows, err := db.c.Query("SELECT id,username FROM users WHERE username LIKE ?", "%"+username+"%")
	log.Println(rows)
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

		p.Photos = Photos
		Profiles = append(Profiles, p)
		log.Println(p.Photos)
	}

	return Profiles, err
}

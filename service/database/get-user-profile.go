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
		var followers int
		var followings int

		Photos, err = db.GetPhotosbyId(int(p.Id))
		if err != nil {
			return Profiles, err
		}

		followers, err = db.GetFollowersbyId(int(p.Id))
		if err != nil {
			return Profiles, err
		}
		followings, err = db.GetFollowingsbyId(int(p.Id))
		if err != nil {
			return Profiles, err
		}

		p.Photos = Photos
		p.N_p = uint64(len(p.Photos))
		p.N_followers = uint64(followers)
		p.N_followings = uint64(followings)
		Profiles = append(Profiles, p)

	}
	if err = rows.Err(); err != nil {
		return Profiles, err
	}
	rows.Close()

	return Profiles, err
}

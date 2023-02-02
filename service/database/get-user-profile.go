package database

import "sort"

func (db *appdbimpl) GetUserProfile(username string, tk int) ([]Profile, error) {
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
		idx, err := db.c.Query("SELECT id FROM users WHERE token=?", tk)
		if err != nil {
			return Profiles, err
		}
		var id int

		for idx.Next() {
			err = idx.Scan(&id)
			if err != nil {
				return Profiles, err
			}
		}

		var count int
		err = db.c.QueryRow("SELECT COUNT(*) FROM banned WHERE idu=? AND idp=?", int(p.Id), id).Scan(&count)
		if err != nil {
			return Profiles, err
		}
		if count == 0 {

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
			sort.Slice(Photos, func(i, j int) bool {
				return Photos[i].Data > Photos[j].Data
			})
			p.N_p = uint64(len(p.Photos))
			p.N_followers = uint64(followers)
			p.N_followings = uint64(followings)
			Profiles = append(Profiles, p)
		}

	}
	if rows.Err() != nil {
		return nil, err
	}
	rows.Close()

	return Profiles, err
}

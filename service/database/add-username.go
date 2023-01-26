package database

func (db *appdbimpl) AddUsername(u User) (User, error) {
	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM users WHERE username=?", u.Username).Scan(&count)
	if err != nil {
		return User{}, err
	}

	if count > 0 {

		err := db.c.QueryRow("SELECT id, username FROM users WHERE username=?", u.Username).Scan(&u.Id, &u.Username)
		if err != nil {
			return User{}, err
		}
		return u, nil
	}

	res, err := db.c.Exec("INSERT INTO users(username) VALUES(?) ", u.Username)
	if err != nil {
		return User{}, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return u, err
	}

	u.Id = uint64(lastInsertID)
	return u, nil
}

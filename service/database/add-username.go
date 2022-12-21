package database

func (db *appdbimpl) AddUsername(u User) (User, error) {
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

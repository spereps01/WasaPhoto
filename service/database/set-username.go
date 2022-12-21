package database

func (db *appdbimpl) SetUsername(u User) (string, error) {
	_, err := db.c.Exec("UPDATE users SET username = ? WHERE id = ?", u.Username, u.Id)
	if err != nil {
		return u.Username, err
	}
	return u.Username, nil

}

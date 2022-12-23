package database

func (db *appdbimpl) SetUsername(u User) (string, error) {

	var count1 int
	err := db.c.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", u.Id).Scan(&count1)
	if err != nil {
		return "", err
	}

	if count1 > 0 {
		_, err := db.c.Exec("UPDATE users SET username = ? WHERE id = ?", u.Username, u.Id)
		if err != nil {
			return u.Username, err
		}
	} else {
		return "utente non presente", nil
	}

	return u.Username, nil

}

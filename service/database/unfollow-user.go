package database

func (db *appdbimpl) UnfollowUser(id1 int, id2 int) (string, error) {

	var count1 int
	err := db.c.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", id1).Scan(&count1)
	if err != nil {
		return "", err
	}

	var count2 int
	err = db.c.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", id2).Scan(&count2)
	if err != nil {
		return "", err
	}

	if count1 > 0 && count2 > 0 {
		_, err := db.c.Exec("DELETE FROM follow WHERE id1=? AND id2=?", id1, id2)
		if err != nil {
			return "", err
		}
	} else {
		return "utente non presente", nil
	}

	return "utente rimosso dai followers", nil
}

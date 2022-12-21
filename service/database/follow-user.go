package database

func (db *appdbimpl) FollowUser(id1 int, id2 int) (string, error) {

	_, err := db.c.Exec("INSERT INTO follow(id1,id2) VALUES(?,?) ", id1, id2)
	if err != nil {
		return "", err
	}

	return "utente seguito correttamente", nil
}

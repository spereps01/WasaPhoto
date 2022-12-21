package database

func (db *appdbimpl) UnfollowUser(id1 int, id2 int) (string, error) {

	_, err := db.c.Exec("DELETE FROM follow WHERE id1=? AND id2=?", id1, id2)
	if err != nil {
		return "", err
	}
	return "utente rimosso dai followers", nil
}

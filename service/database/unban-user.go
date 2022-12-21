package database

func (db *appdbimpl) UnbanUser(idu int, idb int) (string, error) {

	_, err := db.c.Exec("DELETE FROM banned WHERE idu=? AND idp=?", idu, idb)
	if err != nil {
		return "", err
	}
	return "utente rimosso dai bannati", nil
}

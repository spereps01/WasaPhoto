package database

func (db *appdbimpl) BanUser(idu int, idb int) (string, error) {

	_, err := db.c.Exec("INSERT INTO banned(idu,idp) VALUES(?,?) ", idu, idb)
	if err != nil {
		return "", err
	}

	return "utente bannato correttamente", nil
}

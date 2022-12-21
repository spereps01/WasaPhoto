package database

func (db *appdbimpl) AddComment(c Comment) (string, error) {
	_, err := db.c.Exec("INSERT INTO comments(id_ph,comm,owner_id) VALUES(?,?,?) ", c.Id_photo, c.Comment, c.Owner_id)
	if err != nil {
		return "", err
	}
	return "commento aggiunto correttamente", nil
}

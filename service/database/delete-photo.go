package database

func (db *appdbimpl) DeletePhoto(id_photo int, user_id int) (string, error) {
	_, err := db.c.Exec("DELETE FROM photos WHERE id_photo=? AND user_id=?", id_photo, user_id)
	if err != nil {
		return "", err
	}
	return "foto eliminata correttamente", nil
}

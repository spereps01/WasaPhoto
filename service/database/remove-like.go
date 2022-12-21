package database

func (db *appdbimpl) RemoveLike(owner_id int, id_photo int) (string, error) {
	_, err := db.c.Exec("DELETE FROM likes WHERE owner_id=? AND id_photo=?", owner_id, id_photo)
	if err != nil {
		return "", err
	}
	return "like rimosso correttamente", nil
}

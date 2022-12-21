package database

func (db *appdbimpl) AddLike(owner_id int, id_photo int) (string, error) {

	_, err := db.c.Exec("INSERT INTO likes(id_photo,owner_id) VALUES(?,?) ", id_photo, owner_id)
	if err != nil {
		return "", err
	}
	return "like aggiunto correttamente", nil
}

package database

func (db *appdbimpl) AddLike(owner_id int, id_photo int) (string, error) {
	var count1 int
	err := db.c.QueryRow("SELECT COUNT(*) FROM photos WHERE id_photo = ?", id_photo).Scan(&count1)
	if err != nil {
		return "", err
	}

	if count1 > 0 {
		_, err := db.c.Exec("INSERT INTO likes(id_photo,owner_id) VALUES(?,?) ", id_photo, owner_id)
		if err != nil {
			return "", err
		}
	} else {
		return "foto non presente", nil
	}

	return "like aggiunto correttamente", nil
}

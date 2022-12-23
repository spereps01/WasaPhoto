package database

func (db *appdbimpl) AddComment(c Comment) (string, error) {

	var count1 int
	err := db.c.QueryRow("SELECT COUNT(*) FROM photos WHERE id_photo = ?", c.Id_photo).Scan(&count1)
	if err != nil {
		return "", err
	}

	if count1 > 0 {
		_, err := db.c.Exec("INSERT INTO comments(id_ph,comm,owner_id) VALUES(?,?,?) ", c.Id_photo, c.Comment, c.Owner_id)
		if err != nil {
			return "", err
		}
	} else {
		return "foto non presente", nil
	}

	return "commento aggiunto correttamente", nil
}

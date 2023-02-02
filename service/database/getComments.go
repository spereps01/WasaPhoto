package database

func (db *appdbimpl) GetComments(photoid int) ([]Comment, error) {
	var set []Comment
	rows, err := db.c.Query(`SELECT comment_id,id_ph,comm,owner_id FROM comments WHERE id_ph = ?`, photoid)
	if err != nil {
		return set, err
	}
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		var c Comment
		err = rows.Scan(&c.Comment_id, &c.Id_photo, &c.Comment, &c.Owner_id)
		if err != nil {
			return nil, err
		}
		c.Owner_us = db.GetUsernamebyId(int(c.Owner_id))
		set = append(set, c)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return set, nil
}

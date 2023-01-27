package database

func (db *appdbimpl) GetMyStream(tk int) ([]Photo, error) {

	var stream []Photo
	var array []int
	var a int
	var id int

	idx, err := db.c.Query("SELECT id FROM users WHERE token=?", tk)
	if err != nil {
		return stream, err
	}

	for idx.Next() {
		err = idx.Scan(&id)
		if err != nil {
			return stream, err
		}
	}
	if err = idx.Err(); err != nil {
		return stream, err
	}

	foll, err := db.c.Query("SELECT id2 FROM follow WHERE id1=?", id)
	if err != nil {
		return stream, err
	}
	for foll.Next() {

		err = foll.Scan(&a)
		if err != nil {
			return stream, err
		}
		array = append(array, a)

	}

	for _, v := range array {
		var Photos []Photo
		Photos, err = db.GetPhotosbyId(v)

		if err != nil {
			return Photos, err
		}
		stream = append(stream, Photos...)

	}
	foll.Close()

	return stream, err
}

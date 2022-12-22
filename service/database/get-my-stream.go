package database

import "log"

func (db *appdbimpl) GetMyStream(id int) ([]Photo, error) {
	var stream []Photo
	var array []int
	var a int
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
	if err = foll.Err(); err != nil {
		return stream, err
	}

	for _, v := range array {
		var Photos []Photo
		Photos, err = db.GetPhotosbyId(v)
		log.Println(Photos)
		if err != nil {
			return Photos, err
		}
		stream = append(stream, Photos...)

	}

	return stream, err
}

package database

import "log"

func (db *appdbimpl) GetMyStream(id int) ([]Photo, error) {
	var stream []Photo
	var array []int
	var a int
	foll, err := db.c.Query("SELECT id2 FROM follow WHERE id1=?", id)
	for foll.Next() {

		err = foll.Scan(&a)
		array = append(array, a)

	}

	for _, v := range array {
		var Photos []Photo
		Photos, err = db.GetPhotosbyId(v)
		log.Println(Photos)
		if err != nil {
			return Photos, err
		}
		for _, p := range Photos {
			stream = append(stream, p)
		}

	}

	return stream, err
}

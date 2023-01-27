package database

import (
	"math/rand"
	"time"
)

func (db *appdbimpl) AddUsername(u User) (User, error) {

	rand.Seed(time.Now().UnixNano())
	min := 10000
	max := 99999
	token := rand.Intn(max-min+1) + min
	u.Token = uint64(token)

	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM users WHERE username=?", u.Username).Scan(&count)
	if err != nil {
		return User{}, err
	}

	if count > 0 {

		err := db.c.QueryRow("SELECT id, username,token FROM users WHERE username=?", u.Username).Scan(&u.Id, &u.Username, &u.Token)
		if err != nil {
			return User{}, err
		}

		return u, nil
	}

	res, err := db.c.Exec("INSERT INTO users(username,token) VALUES(?,?) ", u.Username, u.Token)
	if err != nil {
		return User{}, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return u, err
	}

	u.Id = uint64(lastInsertID)

	return u, nil
}

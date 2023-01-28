package api

import (
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/database"
)

const (
	FountainStatusGood   string = "good"
	FountainStatusFaulty string = "faulty"
)

// Fountain struct represent a fountain in every data exchange with the external world via REST API. JSON tags have been
// added to the struct to conform to the OpenAPI specifications regarding JSON key names.
// Note: there is a similar struct in the database package. See Fountain.FromDatabase (below) to understand why.

type Users struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Token    uint64 `json:"token"`
}
type Comment struct {
	Comment_id uint64
	Id_photo   uint64
	Comment    string
	Owner_id   uint64
}
type Like struct {
	Id_photo uint64
	Owner_id uint64
}
type Photo struct {
	Id_photo uint64
	User_id  uint64
	Data     string
	Photo    []byte
}

// FromDatabase populates the struct with data from the database, overwriting all values.
// You might think this is code duplication, which is correct. However, it's "good" code duplication because it allows
// us to uncouple the database and API packages.
// Suppose we were using the "database.Fountain" struct inside the API package; in that case, we were forced to conform
// either the API specifications to the database package or the other way around. However, very often, the database
// structure is different from the structure of the REST API.
// Also, in this way the database package is freely usable by other packages without the assumption that structs from
// the database should somehow be JSON-serializable (or, in general, serializable).
func (u *Users) FromDatabase(user database.User) {
	u.ID = user.Id
	u.Username = user.Username
	u.Token = user.Token

}

// ToDatabase returns the fountain in a database-compatible representation
func (u *Users) ToDatabase() database.User {
	return database.User{
		Id:       u.ID,
		Username: u.Username,
		Token:    u.Token,
	}
}

func (p *Photo) ToDatabase_p() database.Photo {
	return database.Photo{
		User_id: p.User_id,
		Data:    p.Data,
		Photo:   p.Photo,
	}
}

func (c *Comment) ToDatabase_c() database.Comment {
	return database.Comment{
		Id_photo: c.Id_photo,
		Comment:  c.Comment,
		Owner_id: c.Owner_id,
	}
}

func (p *Photo) FromDatabase_p(photo database.Photo) {
	p.Id_photo = photo.Id_photo
	p.User_id = photo.User_id
	p.Data = photo.Data
	p.Photo = photo.Photo

}

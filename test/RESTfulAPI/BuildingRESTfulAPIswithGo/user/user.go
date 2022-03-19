package user

import (
	"errors"

	"github.com/asdine/storm/v3"
	"gopkg.in/mgo.v2/bson"
)

//user data
type User struct {
	ID   bson.ObjectId `json:"id" storm:"id"`
	Name string        `json:"name"`
	Role string        `json:"role"`
}

const (
	dbPath = "user.db"
)

//errors
var (
	ErrRecordInvalid = errors.New("record is invalid")
)

// all retrives all users from the database
func All() ([]User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	//main part of the function

	//create the users list
	users := []User{}
	err = db.All(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// return only a single user info
func One(id bson.ObjectId) (*User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	u := new(User)
	err = db.One("ID", id, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func Delete(id bson.ObjectId) error {
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	u := new(User)
	err = db.One("ID", id, u)
	if err != nil {
		return err
	}

	return db.DeleteStruct(u)
}

func (u *User) Save() error {
	if err := u.validate(); err != nil {
		return err
	}
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Save(u)
}

// validate makes sure that the record contains valid data
func (u *User) validate() error {
	if u.Name == "" {
		return ErrRecordInvalid
	}
	return nil
}

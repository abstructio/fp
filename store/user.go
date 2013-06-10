package store

import (
	"crypto/sha256"
	"fmt"
	"labix.org/v2/mgo/bson"
)

//hash a string with sha256, return hex string
func sha(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	r := h.Sum(nil)
	return fmt.Sprintf("%x", r)
}

//Create a new password hash
func NewPasswordHash(raw string) passwordHash {
	return passwordHash(sha(raw))
}

//Represents a password
type passwordHash string

//User represents a user in fp
type User struct {
	Id       bson.ObjectId `bson:"_id"`
	Username string
	Surename string
	Lastname string
	Password passwordHash
}

//Create a new user, shorthand for manual usersetup
func NewUser(username, surename, lastname, password string) *User {
	return &User{
		Username: username,
		Surename: surename,
		Lastname: lastname,
		Password: NewPasswordHash(password),
	}
}

//Insert or Update a provided User.
//Returns true if updated.
//If an error occurs, its returned.
//Attention, this function doesn't check if the username is already taken!!!
func InsertOrUpdateUser(u *User) (bool, error) {
	session := connect()
	defer session.Close()

	if u.Id.Hex() == "" {
		u.Id = bson.NewObjectId()
	}

	info, err := session.DB("").C("user").UpsertId(u.Id, u)
	if err != nil {
		panic(err)
	}

	return info.Updated == 0, err

}

//Returns an error if no user is found.
func UserById(id bson.ObjectId) (*User, error) {
	session := connect()
	defer session.Close()

	result := User{}

	c := session.DB("").C("user")

	err := c.Find(bson.M{"_id": id}).One(&result)
	return &result, err
}

//Check if username is already taken or not, primary use for registry
func IsUsernameAviable(username string) bool {
	session := connect()
	defer session.Close()

	err := session.DB("").C("user").Find(bson.M{"username": username})

	return err == nil
}

//Check if username and password combination is right
func CheckUsernamePassword(username, password string) bool {
	session := connect()
	defer session.Close()

	count, err := session.DB("").C("user").Find(bson.M{"username": username,
		"password": NewPasswordHash(password)}).Count()
	if err != nil {
		panic(err)
	}

	return count == 1 && err == nil

}

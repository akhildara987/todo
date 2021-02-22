package user

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

// User :
type User struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"Lastname"`
	Password  string `json:"password,omitempty"`
	Email     string `json:"email"`
}

//LoginParams :
type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

func (u *LoginParams) fromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}
func (u *User) fromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

//Getunique :
func Getunique() string {

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	uuid := fmt.Sprintf("%x", b[0:8])

	return uuid
}

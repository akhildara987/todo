package todo

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

// Todo :
type Todo struct {
	ID           string `json:"ID"`
	TodoID       string `json:"todoID"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	EndDate      string `json:"enddate"`
	TodoPriority string `json:"todopriority'`
	Iscompleted  string `json:"is_completed"`
	TodoType     string `json:todotype`
}

func (t *Todo) fromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(t)
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

package user

import (
	"encoding/json"
	"log"
	"net/http"
)

// RequestPost :
func (s *Services) RequestPost(w http.ResponseWriter,
	r *http.Request) {

	usr := &User{}
	err := usr.fromJSON(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err = s.ls.AddUser(usr)
	if err != nil {
		http.Error(w, "error in db", http.StatusBadRequest)
		return
	}

	success := struct {
		Message string `json:"message"`
	}{
		Message: "Account Created",
	}

	jsonStr, _ := json.Marshal(success)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonStr)

}

// RequestGet :
func (s *Services) RequestGet(w http.ResponseWriter,
	r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "missing parameter :- id", http.StatusBadRequest)
		return
	}

	usr, err := s.ls.GetUser(id)
	if err != nil {
		http.Error(w, "error in db", http.StatusInternalServerError)
		return
	}

	jsonStr, _ := json.Marshal(usr)
	log.Printf("%s\n", jsonStr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonStr)
}

// RequestPut :
func (s *Services) RequestPut(w http.ResponseWriter,
	r *http.Request) {

	usr := &User{}
	err := usr.fromJSON(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err = s.ls.UpdateUser(usr)
	if err != nil {
		http.Error(w, "Error in db", http.StatusBadRequest)
		return
	}

	success := struct {
		Message string `json:"message"`
	}{
		Message: "account updated",
	}

	jsonStr, _ := json.Marshal(success)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonStr)

}

// RequestDelete :
func (s *Services) RequestDelete(w http.ResponseWriter,
	r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "missing parameter :- id", http.StatusBadRequest)
		return
	}

	err := s.ls.deleteUser(id)
	if err != nil {
		http.Error(w, "error in db", http.StatusInternalServerError)
		return
	}

	success := struct {
		Message string `json:"message"`
	}{
		Message: "account deleted",
	}

	jsonStr, _ := json.Marshal(success)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonStr)

}

package todo

import (
	"encoding/json"
	"log"
	"net/http"
)

// RequestPost :
func (s *Services) RequestPost(w http.ResponseWriter,
	r *http.Request) {
	todo := &Todo{}
	err := todo.fromJSON(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	todo.ID = r.URL.Query().Get("id")
	if todo.ID == "" {
		http.Error(w, "missing parameter :- id", http.StatusBadRequest)
		return
	}

	err = s.ls.Addtodo(todo)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error in db", http.StatusInternalServerError)
		return
	}
	success := struct {
		Message string `json:"message"`
	}{
		Message: "Todo Created",
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
	todoID := r.URL.Query().Get("todoid")

	if id == "" {
		http.Error(w, "missing parameter :- id", http.StatusBadRequest)
		return
	}

	// if todoID == "" {
	// 	http.Error(w, "missing parameter :- todoID", http.StatusBadRequest)
	// 	return
	// }

	list, err := s.ls.Gettodo(id, todoID)
	if err != nil {
		http.Error(w, "error in db"+err.Error(), http.StatusInternalServerError)
		return
	}

	jsonStr, _ := json.Marshal(list)
	log.Printf("%s\n", jsonStr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonStr)

}

// RequestPut :
func (s *Services) RequestPut(w http.ResponseWriter,
	r *http.Request) {

	todo := &Todo{}
	err := todo.fromJSON(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	todo.ID = r.URL.Query().Get("id")
	if todo.ID == "" {
		http.Error(w, "missing parameter :- id", http.StatusBadRequest)
		return
	}

	if todo.TodoID == "" {
		http.Error(w, "missing parameter :- todoID", http.StatusBadRequest)
		return
	}

	err = s.ls.Updatetodo(todo)
	if err != nil {
		http.Error(w, "Error in db"+err.Error(), http.StatusBadRequest)
		return
	}
	success := struct {
		Message string `json:"message"`
	}{
		Message: "Todo updated",
	}

	jsonStr, _ := json.Marshal(success)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonStr)
}

// RequestDelete :
func (s *Services) RequestDelete(w http.ResponseWriter,
	r *http.Request) {

	id := r.URL.Query().Get("id")
	todoID := r.URL.Query().Get("todoid")

	if id == "" {
		http.Error(w, "missing parameter :- id", http.StatusBadRequest)
		return
	}

	if todoID == "" {
		http.Error(w, "missing parameter :- todoID", http.StatusBadRequest)
		return
	}

	err := s.ls.deletetodo(todoID, id)
	if err != nil {
		http.Error(w, "Error in db"+err.Error(), http.StatusInternalServerError)
		return
	}

	success := struct {
		Message string `json:"message"`
	}{
		Message: "todo deleted",
	}

	jsonStr, _ := json.Marshal(success)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonStr)

}

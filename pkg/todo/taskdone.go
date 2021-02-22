package todo

import (
	"encoding/json"
	"errors"
	"net/http"
)

// TodoCompleted :
func (s *Services) TodoCompleted(w http.ResponseWriter,
	r *http.Request) {

	todo := &Todo{}
	// err := todo.fromJSON(r.Body)
	// defer r.Body.Close()
	// if err != nil {
	// 	http.Error(w, "bad request", http.StatusBadRequest)
	// 	return
	// }
	todo.ID = r.URL.Query().Get("id")
	todo.TodoID = r.URL.Query().Get("todoid")

	if todo.ID == "" {
		http.Error(w, "missing parameter :- id", http.StatusBadRequest)
		return
	}

	if todo.TodoID == "" {
		http.Error(w, "missing parameter :- todoID", http.StatusBadRequest)
		return
	}

	err := s.ls.completedTodo(todo)
	if err != nil {
		http.Error(w, "Error in db"+err.Error(), http.StatusInternalServerError)
		return
	}
	success := struct {
		Message string `json:"message"`
	}{
		Message: "Todo updated to completed",
	}

	jsonStr, _ := json.Marshal(success)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonStr)
}

// UpdateUser :
func (ls *localServices) completedTodo(todo *Todo) error {
	var paramas []interface{}
	query := " UPDATE todo set"
	query += " iscompleted=?"
	query += " WHERE todoID=? and id =?"
	paramas = append(paramas, "true")
	paramas = append(paramas, todo.TodoID)
	paramas = append(paramas, todo.ID)

	pre, err := ls.db.Prepare(query)
	if err != nil {
		return err
	}

	out, err := pre.Exec(paramas...)

	if rows, err := out.RowsAffected(); rows == 1 && err == nil {
		return nil
	} else if err != nil {
		return err
	}

	return errors.New("unable to process this request")
}

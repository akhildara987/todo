package restservice

import (
	"net/http"
	"todolist/pkg/todo"
	"todolist/pkg/user"
)

// RestHandler :
func RestHandler(u user.RestServices, t todo.RestServices) *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/api/user", u.HandleRequests)
	mux.HandleFunc("/api/todo", t.HandleRequests)
	mux.HandleFunc("/api/todo/completed", t.TodoCompleted)

	mux.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("../web"))))

	mux.HandleFunc("/api/login", u.Login)

	return mux

}

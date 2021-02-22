package todo

import (
	"database/sql"
	"net/http"
)

// RestServices :
type RestServices interface {
	HandleRequests(w http.ResponseWriter, r *http.Request)
	TodoCompleted(w http.ResponseWriter, r *http.Request)
}

// Services :
type Services struct {
	ls localServices
}

// localServices :
type localServices struct {
	db *sql.DB
}

// NewServices :
func NewServices(db *sql.DB) RestServices {
	return &Services{
		ls: localServices{
			db: db,
		},
	}
}

//HandleRequests :
func (s *Services) HandleRequests(w http.ResponseWriter,
	r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		s.RequestPost(w, r)
	case http.MethodGet:
		s.RequestGet(w, r)
	case http.MethodPut:
		s.RequestPut(w, r)
	case http.MethodDelete:
		s.RequestDelete(w, r)
	default:
		http.Error(w, "Method Not Allowed",
			http.StatusMethodNotAllowed)
	}

}

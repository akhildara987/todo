package user

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

// Login :
func (s *Services) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed",
			http.StatusMethodNotAllowed)
		return
	}

	usr := &LoginParams{}

	err := usr.fromJSON(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if usr.Email == "" {
		http.Error(w, "missing parameter :- email", http.StatusBadRequest)
		return

	}
	if usr.Password == "" {
		http.Error(w, "missing parameter :- password", http.StatusBadRequest)
		return
	}
	usrData, err := s.ls.GetUserLogin(usr.Email)
	if err != nil {
		http.Error(w, "error in db", http.StatusInternalServerError)
		return
	}
	if usrData.Password == usr.Password {
		usrData.Password = ""
		rdurl := struct {
			RediretcUrl string `json:"redirecturl"`
			ID          string `json:"id"`
		}{
			RediretcUrl: "/web/todo.html",
			ID:          usrData.ID,
		}
		jsonStr, _ := json.Marshal(rdurl)
		log.Printf("%s\n", jsonStr)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonStr)
		// http.Redirect(w, r, "/api/todo?id="+usrData.ID, http.StatusFound)
		return

	}
	usrData.Password = ""

	// http.Redirect(w, r, "/api/todo?id="+usrData.ID, 301)
	http.Error(w, "invalid password or email", http.StatusInternalServerError)

}

func (ls *localServices) GetUserLogin(email string) (*User, error) {
	usr := &User{}
	query := " SELECT "
	query += " Firstname,lastname,email,id,password "
	query += " from user WHERE email =?"

	sel := ls.db.QueryRow(query, email)

	err := sel.Scan(&usr.Firstname, &usr.Lastname, &usr.Email, &usr.ID, &usr.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return usr, nil
		}
		return usr, err
	}
	return usr, nil
}

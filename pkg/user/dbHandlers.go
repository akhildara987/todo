package user

import (
	"database/sql"
	"errors"
)

func (ls *localServices) AddUser(usr *User) error {
	usr.ID = Getunique()
	query := " INSERT INTO user"
	query += " (id,Firstname,lastname,email,password)"
	query += " VALUES(?,?,?,?,?)"

	pre, err := ls.db.Prepare(query)
	if err != nil {
		// http.Error(w, "error in db", http.StatusInternalServerError)
		return err
	}

	out, err := pre.Exec(usr.ID, usr.Firstname, usr.Lastname,
		usr.Email, usr.Password)

	if rows, err := out.RowsAffected(); rows == 1 && err == nil {
		return nil
	} else if err != nil {
		return err
	}

	return errors.New("some thing went wrong")
}

func (ls *localServices) GetUser(id string) (*User, error) {
	usr := &User{}
	query := " SELECT "
	query += " Firstname,lastname,email "
	query += " from user WHERE id =?"

	sel := ls.db.QueryRow(query, id)

	err := sel.Scan(&usr.Firstname, &usr.Lastname, &usr.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return usr, nil
		}
		return usr, err
	}
	return usr, nil
}

func (ls *localServices) UpdateUser(usr *User) error {
	var params []interface{}
	query := " UPDATE user SET"
	if usr.Lastname != "" {
		query += " lastname=?,"
		params = append(params, usr.Lastname)

	}
	if usr.Firstname != "" {
		query += " firstname=?,"
		params = append(params, usr.Firstname)

	}
	if usr.Email != "" {
		query += " email=?,"
		params = append(params, usr.Email)
	}

	query += " id=? WHERE id =?"
	params = append(params, usr.ID)
	params = append(params, usr.ID)

	pre, err := ls.db.Prepare(query)
	if err != nil {
		// http.Error(w, "error in db", http.StatusInternalServerError)
		return err
	}

	out, err := pre.Exec(params...)

	if rows, err := out.RowsAffected(); rows == 1 && err == nil {
		return nil
	} else if err != nil {
		return err
	}

	return errors.New("some thing went wrong")
}

func (ls *localServices) deleteUser(id string) error {

	query := "DELETE FROM user WHERE id =?"

	out, err := ls.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := out.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 1 {
		return nil
	}

	return errors.New("Nothing to delete")
}

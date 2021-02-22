package todo

import (
	"database/sql"
	"errors"
	"log"
)

func (ls *localServices) Addtodo(todo *Todo) error {
	log.Println("hello")
	todo.TodoID = Getunique()
	query := " INSERT INTO todo"
	query += " (id,todoID,title,description,enddate,todopriority,todotype)"
	query += " VALUES(?,?,?,?,?,?,?)"

	pre, err := ls.db.Prepare(query)
	if err != nil {
		return err
	}
	log.Println("hello 2")

	out, err := pre.Exec(todo.ID, todo.TodoID, todo.Title, todo.Description,
		todo.EndDate, todo.TodoPriority, todo.TodoType)
	log.Println("hello 3")
	if err != nil {
		return err
	}

	if rows, err := out.RowsAffected(); rows == 1 && err == nil {
		return nil
	} else if err != nil {
		return err
	}
	log.Println("hello 4")

	return errors.New("unable to process this request")
}

func (ls *localServices) Gettodo(id, todoID string) (*[]Todo, error) {

	todo := &[]Todo{}

	out := &Todo{}

	var paramas []interface{}
	query := "SELECT "
	query += " id,todoID,title,description,enddate,todopriority,iscompleted,todotype"
	query += " FROM todo WHERE id =? "
	paramas = append(paramas, id)

	if todoID != "" {
		query += "and todoID=?"
		paramas = append(paramas, todoID)

	}

	sel, err := ls.db.Query(query, paramas...)
	if err != nil {
		if err == sql.ErrNoRows {
			return todo, nil
		}
		return todo, err
	}

	for sel.Next() {

		err := sel.Scan(
			&out.ID,
			&out.TodoID,
			&out.Title,
			&out.Description,
			&out.EndDate,
			&out.TodoPriority,
			&out.Iscompleted,
			&out.TodoType,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return todo, nil
			}
			return todo, err
		}

		*todo = append(*todo, *out)
	}
	return todo, nil

}

// UpdateUser :
func (ls *localServices) Updatetodo(todo *Todo) error {
	var paramas []interface{}
	query := " UPDATE todo set"
	if todo.Title != "" {
		query += " title=?,"
		paramas = append(paramas, todo.Title)
	}
	if todo.Description != "" {
		query += " description=?,"

		paramas = append(paramas, todo.Description)
	}
	if todo.EndDate != "" {
		query += " enddate=?,"

		paramas = append(paramas, todo.EndDate)
	}
	if todo.TodoPriority != "" {
		query += " todopriority=?,"
		paramas = append(paramas, todo.TodoPriority)
	}
	if todo.TodoType != "" {
		query += " todotype=?,"
		paramas = append(paramas, todo.TodoType)
	}

	query += " todoID=? WHERE todoID=? and id =?"
	paramas = append(paramas, todo.TodoID)
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

func (ls *localServices) deletetodo(todoID, id string) error {

	query := " DELETE FROM todo where todoID=? and id =? "

	out, err := ls.db.Exec(query, todoID, id)
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

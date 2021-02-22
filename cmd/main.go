package main

import (
	"database/sql"
	"net/http"
	"time"
	"todolist/pkg/restservice"
	"todolist/pkg/todo"
	"todolist/pkg/user"

	_ "github.com/go-sql-driver/mysql"
)

const port = ":800"

// DbSchema :
type DbSchema struct {
	dbEngine   string
	dbUsername string
	dbPassword string
	dbPort     string
	dbName     string
	dbHost     string
}

func main() {
	dbData := DbSchema{
		dbEngine:   "mysql",
		dbUsername: "user",
		dbPassword: "password",
		dbPort:     "3306",
		dbName:     "todo",
		dbHost:     "localhost",
	}
	db := dbData.ConnectDB()

	ss := user.NewServices(db)
	tt := todo.NewServices(db)

	http.ListenAndServe(port, restservice.RestHandler(ss, tt))

}

//ConnectDB :
func (sqldata *DbSchema) ConnectDB() *sql.DB {

	db, err := sql.Open(sqldata.dbEngine,
		sqldata.dbUsername+":"+sqldata.dbPassword+"@tcp("+sqldata.dbHost+":"+
			sqldata.dbPort+")/"+sqldata.dbName)
	if err != nil {
		panic(err.Error())
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(15)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db
}

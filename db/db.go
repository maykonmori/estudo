package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func conectaComBaseDeDados() *sql.DB {
	conexao := "user=postgres dbname=db_estudos password=db_&studos host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

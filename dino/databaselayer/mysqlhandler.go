package databaselayer

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type MYSQLHandler struct {
	*SQLHandler
}

func NewMYSQLHandler(connection string) (*MYSQLHandler, error) {
	db, err := sql.Open("mysql", connection)
	return &MYSQLHandler{
		SQLHandler: &SQLHandler{
			DB: db
		},
	}, err
}
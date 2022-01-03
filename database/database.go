package database

import "database/sql"

var db *sql.DB

func ConnectDatabase() (*sql.DB, error) {


	return db, nil
}
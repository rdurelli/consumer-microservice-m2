package database

import (
	"database/sql"
	"log"
)

type Db struct {
	Db *sql.DB
}

func NewDataBase() Db {
	DB, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/emailqueue?parseTime=true")
	if err != nil {
		log.Fatal("Something went wrong trying to connect to the database ", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Something went wrong trying to ping to the database ")
	}
	return Db{
		Db: DB,
	}
}

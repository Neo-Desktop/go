package main

import (
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var DB *sqlx.DB

func sqlInit() *sqlx.DB {
	log.Println("Opening SQL Connection")
	DB, err := sqlx.Open("mysql", "gouser:gopassword@tcp(127.0.0.1:3309)/goGrounds")

	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	DB.SetMaxIdleConns(1)

	log.Println("SQL Connection opened")
	return DB
}

func sqlClose() {
	DB.Close()
	log.Println("SQL Connection gracefully closed")
}

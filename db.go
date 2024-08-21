package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func ConnectDataBase() *sql.DB {
	senha := os.Getenv("DB_PASS")
	connStr := fmt.Sprint("user=postgres dbname=digport_academy_attendance password=", senha, " host=localhost sslmode=disable")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

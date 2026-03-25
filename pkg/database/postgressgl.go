package database

import (
	"database/sql"
	"fmt"
	"log"
	"mylibary/configs"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnecDB(conn *configs.PostgresSql) *sql.DB {
	host := conn.Host
	port := conn.Port
	user := conn.Username
	dbname := conn.DBname
	password := conn.Password
	sslmode := conn.SSlmode

	connecDB := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, user, dbname, password, sslmode)

	db, err := sql.Open("postgres", connecDB)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect Database Successfully!!!")
	return db

}

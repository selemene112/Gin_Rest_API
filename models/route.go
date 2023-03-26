package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
    host     = "127.0.0.1"
    port     = 5432
    user     = "admin"
    password = "lele123"
    dbname   = "local"
)
var (
db *sql.DB
err error
)


func ConnectDatabase(){


	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " + 
	"password =%s dbname=%s sslmode=disable", 
	host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil{
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}


    // db, err := sql.Open("postgres", "postgres://admin:lele123@host:5432/local?sslmode=disable")
    // if err != nil {
    //     panic(err)
    // }
    // defer db.Close()

    // err = db.Ping()
    // if err != nil {
    //     panic(err)
    // }
    fmt.Println("Connected to PostgreSQL database!")
}
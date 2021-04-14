package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)


func InitDb() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "root:@tcp(127.0.0.1:3306)/UF")
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
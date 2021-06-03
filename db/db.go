package db

import (
	"log"

	_ "database/sql"

	_ "github.com/bmizerany/pq"
	// _ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	// _ "github.com/lib/pq"
)

var Db *sqlx.DB

// func Connect() *sqlx.DB {
// 	db := sqlx.MustConnect("mysql", "root:@tcp(127.0.0.1:3306)/db_belajar_golang")
// 	log.Println(db)
//     return db
// }

// "postgres://postgres:@localhost:5432/postgres?sslmode=disable"

func Connect() *sqlx.DB {
	db := sqlx.MustConnect("postgres", "user=postgres password= dbname=postgres sslmode=disable")
	log.Println(db)
	return db
}

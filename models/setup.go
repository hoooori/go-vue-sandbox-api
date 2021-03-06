package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/boil"
	"log"
	"os"
)

var DB *sql.DB

func ConnectDB() *sql.DB {
	var err error

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	DB, err = sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatalf("failed to connect database: %s", err)
	}

	boil.SetDB(DB)
	return DB
}

package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	connStr := fmt.Sprintf("host=%s port=5432 user=b1user password=b1pass dbname=b1trainer sslmode=disable", host)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open database: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Connected to PostgreSQL!")
}

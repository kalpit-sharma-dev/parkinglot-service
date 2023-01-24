package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	HOST = "kalpitcool2006-gmailcom_my-postgres_1"
	PORT = 5432
)

// ErrNoMatch is returned when we request a row that doesn't exist
var ErrNoMatch = fmt.Errorf("no matching record")

type Database struct {
	Conn *sql.DB
}

func GetDatabaseProvider(username, password, database string) (*sql.DB, error) {
	//db := Database{}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, username, password, database)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return conn, err
	}
	//db.Conn = conn
	err = conn.Ping()
	if err != nil {
		return conn, err
	}
	log.Println("Database connection established")
	return conn, err
}

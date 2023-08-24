package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var (
	Conn *sql.DB
)

func InitializeDB(connStr string) error {
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err := conn.Ping(); err != nil {
		return err
	}

	Conn = conn

	initializeTables()

	return nil
}

func CloseDB() {
	if Conn != nil {
		Conn.Close()
	}
}

func initializeTables() {
	createPeopleTable()
}

func createPeopleTable() {
	Conn.Exec(`CREATE TABLE IF NOT EXISTS people (
		id UUID PRIMARY KEY,
		nickname CHAR(32) NOT NULL,
		name CHAR(100) NOT NULL,
		birthdate DATE NOT NULL,
		stack CHAR(32)[]
	);`)
}

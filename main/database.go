package main

import (
	"database/sql"
)

func Databse() {
	databaseFile := "./rs.db"
	InitDB(databaseFile)
}

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}
	return db
}

func CreateTables(db *sql.DB) {
	// create tables if not exists
	sql_table := `
	CREATE TABLE IF NOT EXISTS invites(
		Id TEXT NOT NULL PRIMARY KEY,
		Name TEXT,
		Email TEXT,
		Description TEXT,
		ProcessedDatetime DATETIME,
		InsertedDatetime DATETIME
	);
	`

	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}

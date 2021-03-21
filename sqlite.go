package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type People struct {
	ID        uint
	Firstname string
	Lastname  string
}

func SQLite() {
	// open database
	DB, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	}

	// create table people
	statement, err := DB.Prepare(`CREATE TABLE IF NOT EXISTS people (
        id INTEGER PRIMARY KEY, 
        firstname TEXT,
        lastname TEXT);`)
	statement.Exec()
	if err != nil {
		panic(err)
	}

	// insert some values
	statement, err = DB.Prepare("INSERT INTO people (firstname, lastname) VALUES (?,?)")
	statement.Exec("saul", "pinales")
	if err != nil {
		panic(err)
	}

	// quering database
	rows, err := DB.Query("SELECT * FROM people")
	if err != nil {
		panic(err)
	}
	var p People
	for rows.Next() {
		rows.Scan(&p.ID, &p.Firstname, &p.Lastname)
		fmt.Println(p)
	}
}

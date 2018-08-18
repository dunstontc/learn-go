package main

import (
	"database/sql"
	"fmt"
	"log"
	"os/user"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// ListTables lists all tables in a given DB.
func ListTables(filepath string) {
	db := InitDB(expand(filepath))
	rows, _ := db.Query("SELECT name FROM sqlite_master WHERE type='table'")
	var name string
	for rows.Next() {
		rows.Scan(&name)
		fmt.Println(name)
	}
	rows.Close()
	db.Close()
}

func ListColumns(filepath string) {
	db := InitDB(expand(filepath))
	rows, _ := db.Query("SELECT * FROM sqlite_master WHERE type='table'")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	fmt.Println(rows.Columns())
	rows.Close()
	db.Close()
}

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", expand(filepath))
	if err != nil {
		log.Fatal(err.Error())
	}
	if db == nil {
		log.Fatal("db nil")
	}
	return db
}

/* SEE ALSO: https://github.com/mitchellh/go-homedir */
func expand(path string) string {
	if len(path) == 0 || path[0] != '~' {
		return path
	}

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}
	return filepath.Join(usr.HomeDir, path[1:])
}

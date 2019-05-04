package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS test (key TEXT PRIMARY KEY, value TEXT)")
	statement.Exec()

	statement, _ = db.Prepare("INSERT INTO TEST (KEY,VALUE) VALUES (?,?)")
	statement.Exec("golang", "awesome")

	row := db.QueryRow("select key,value FROM TEST WHERE key = ?", "golang")
	var value string
	var key string
	err = row.Scan(&key, &value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(key + " is " + value)
}

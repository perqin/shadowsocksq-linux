package sqlite

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Open() {
    sqlite3Db, err := sql.Open("sqlite3", "/tmp/ssq.db")
    if err == nil {
        // TODO: Create table
        db.Exec("CREATE TABLE IF NOT EXISTS profiles (id INTEGER PRIMARY KEY AUTO INCREMENT)")
        db = sqlite3Db
    }
}

func Close() {
    db.Close()
}

func GetDb() *sql.DB {
    return db
}

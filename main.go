package main

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/mattn/go-sqlite3"
)

func initializeDatabase() error {
    // Abra a conexão com o banco de dados SQLite3
    db, err := sql.Open("sqlite3", "./test.db")
    if err != nil {
        return err
    }
    defer db.Close()

    // Crie uma tabela de exemplo
    sqlStmt := `
    CREATE TABLE IF NOT EXISTS example (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT
    );
    `
    _, err = db.Exec(sqlStmt)
    if err != nil {
        return err
    }

    fmt.Println("Database initialized and ready.")
    return nil
}

func main() {
    if err := initializeDatabase(); err != nil {
        log.Fatal(err)
    }
}

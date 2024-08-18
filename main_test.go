package main

import (
    "database/sql"
    "testing"
    _ "github.com/mattn/go-sqlite3"
)

func TestInitializeDatabase(t *testing.T) {
    // Chame a função que inicializa o banco de dados
    if err := initializeDatabase(); err != nil {
        t.Fatalf("Failed to initialize database: %v", err)
    }

    // Verifique se o banco de dados foi criado e a tabela está presente
    db, err := sql.Open("sqlite3", "./test.db")
    if err != nil {
        t.Fatalf("Failed to open database: %v", err)
    }
    defer db.Close()

    // Verifique se a tabela 'example' existe
    row := db.QueryRow("SELECT count(*) FROM sqlite_master WHERE type='table' AND name='example'")
    var count int
    if err := row.Scan(&count); err != nil {
        t.Fatalf("Failed to query table existence: %v", err)
    }

    if count != 1 {
        t.Fatalf("Table 'example' does not exist")
    }
}

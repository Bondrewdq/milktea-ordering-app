package config

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
    db, err := sql.Open("sqlite3", "./milktea.db")
    if err != nil {
        log.Fatal("Failed to connect database:", err)
    }

    // 创建表（示例：奶茶产品）
    createTableSQL := `
    CREATE TABLE IF NOT EXISTS teas (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        price REAL NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );`
    
    if _, err := db.Exec(createTableSQL); err != nil {
        log.Fatal("Failed to create table:", err)
    }

    log.Println("Database initialized successfully")
    return db
}
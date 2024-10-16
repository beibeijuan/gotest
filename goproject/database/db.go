// database/db.go
package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Connect establishes a connection to the database
func Connect() {
	var err error
	// Update the DSN with your MySQL credentials
	DB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("无法连接到数据库: %v", err)
	}
}

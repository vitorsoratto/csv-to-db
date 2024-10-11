package config

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBPath string = "./db/fitness.db"

func InitSqlite() (*gorm.DB, error) {
	// First check if the database file exists
	_, err := os.Stat(DBPath)
	if os.IsNotExist(err) {
		fmt.Println("Database file not found")
	}

  // Open the connection to the database
	db, err := gorm.Open(sqlite.Open(DBPath), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error opening sqlite database: %v", err)
		return nil, err
	}

	return db, nil
}

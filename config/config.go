package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
)

func InitDatabase() (*gorm.DB, error) {
	var err error
	db, err = InitSqlite()
	if err != nil {
		return nil, err
	}

	return db, nil
}

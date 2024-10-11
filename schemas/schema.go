package schemas

import (
	"github.com/vitorsoratto/csv-to-db/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Schema interface {
	Insert(data interface{}) error
}

func InitSchema() error {
	confDB, err := config.InitDatabase()
	if err != nil {
		return err
	}
	db = confDB
	
	return nil
}


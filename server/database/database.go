package database

import (
	"github.com/YUDAI-HIZU/gin-api/server/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GormConnect() {
	db, err := gorm.Open("mysql", config.DatabaseURL)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	db.LogMode(true)
}

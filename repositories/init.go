package repositories

import (
	"gorm.io/gorm"
)

type connections struct {
	DB *gorm.DB
}

var Connections connections

func Init() {
	initDB()
}

package database

import (
	"gorm.io/gorm"
)

const (
	DB_NAME = "moneycontrol"
	DB_HOST = "localhost"
	DB_USER = "postgres"
	DB_PASS = "postgres"
	DB_PORT = "5432"
)

var instance *gorm.DB

func ConnectSingleton() {

}

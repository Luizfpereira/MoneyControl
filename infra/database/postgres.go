package database

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_NAME = "moneycontrol"
	DB_HOST = "db"
	DB_USER = "postgres"
	DB_PASS = "postgres"
	DB_PORT = "5432"
)

var instance *gorm.DB
var connectionString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORT)
var lock = &sync.Mutex{}

func ConnectSingleton() *gorm.DB {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			log.Println("Creating database instance...")
			var err error
			instance, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
			if err != nil {
				log.Panic("Failed to connect to database!")
			}
		} else {
			log.Println("Connected to database!")
		}
	} else {
		log.Println("Connected to database!")
	}
	return instance
}

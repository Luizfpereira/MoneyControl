package database

import (
	"log"

	"gorm.io/gorm"
)

func Migrate(instance *gorm.DB) {
	if err := instance.AutoMigrate(&User{}, &Transaction{}); err != nil {
		log.Fatalln("Could not migrate models to database!")
	}
	log.Println("Database migration completed!")
}

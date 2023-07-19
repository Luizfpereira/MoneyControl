package database

import (
	"log"

	"gorm.io/gorm"
)

func Migrate(instance *gorm.DB) {
	instance.AutoMigrate(&User{}, &Transaction{})
	// inserir FK em transaction na migração
	log.Println("Database migration completed!")
}

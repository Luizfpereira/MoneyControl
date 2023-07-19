package database

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	LastName string
	Email    string
	Password string
}

type Transaction struct {
	gorm.Model
	Value       float64
	Description string
	Category    string
	Date        time.Time
	UserID      uint
}

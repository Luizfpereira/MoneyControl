package database

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Value       float64 `gorm:"not null"`
	Description string  `gorm:"size:500"`
	Category    string  `gorm:"size:200"`
	Date        time.Time
	UserID      uint
}

// definir tamanho de colunas e campos not null
type User struct {
	gorm.Model
	Name         string `gorm:"size:200"`
	LastName     string `gorm:"size:200"`
	Email        string `gorm:"unique; not null"`
	Password     string `gorm:"not null"`
	Transactions []Transaction
}

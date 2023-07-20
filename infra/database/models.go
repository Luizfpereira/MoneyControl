package database

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Value       float64
	Description string
	Category    string
	Date        time.Time
	UserID      uint
}

// definir tamanho de colunas e campos not null
type User struct {
	gorm.Model
	Name         string
	LastName     string
	Email        string
	Password     string
	Transactions []Transaction
}

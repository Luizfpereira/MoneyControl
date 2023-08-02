package entity

import (
	"errors"
	"time"
)

var (
	ErrInvalidUserID      = errors.New("invalid user_id")
	ErrInvalidDescription = errors.New("invalid description")
	ErrInvalidCategory    = errors.New("invalid category")
	ErrInvalidDate        = errors.New("invalid date")
)

type Transaction struct {
	ID          uint      `json:"id"`
	Value       float64   `json:"value"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Date        time.Time `json:"date"`
	UserID      uint      `json:"user_id"`
}

func NewTransaction(value float64, description, category string, date time.Time, userID uint) (*Transaction, error) {
	transaction := &Transaction{
		Value:       value,
		Description: description,
		Category:    category,
		Date:        date,
		UserID:      userID,
	}
	if err := transaction.Validate(); err != nil {
		return nil, err
	}
	return transaction, nil
}

func (t *Transaction) Validate() error {
	if t.UserID < 1 {
		return ErrInvalidUserID
	}

	if t.Description == "" {
		return ErrInvalidDescription
	}

	if t.Category == "" {
		return ErrInvalidCategory
	}

	if t.Date.IsZero() {
		return ErrInvalidDate
	}
	return nil
}

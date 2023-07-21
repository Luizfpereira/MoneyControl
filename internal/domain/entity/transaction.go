package entity

import (
	"errors"
	"time"
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
	if err := transaction.Validade(); err != nil {
		return nil, err
	}
	return transaction, nil
}

func (t *Transaction) Validade() error {
	if t.UserID < 1 {
		return errors.New("invalid user_id")
	}

	if t.Description == "" {
		return errors.New("invalid description")
	}

	if t.Category == "" {
		return errors.New("invalid description")
	}

	if t.Date.IsZero() {
		return errors.New("invalid date")
	}
	return nil
}

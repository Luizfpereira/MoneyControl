package entity

import (
	"errors"
	"time"
)

var (
	errInvalidUserID      = errors.New("invalid user_id")
	errInvalidDescription = errors.New("invalid description")
	errInvalidCategory    = errors.New("invalid category")
	errInvalidDate        = errors.New("invalid date")
)

//Transaction defines a struct transaction made by a user
type Transaction struct {
	ID          uint      `json:"id"`
	Value       float64   `json:"value"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Date        time.Time `json:"date"`
	UserID      uint      `json:"user_id"`
}

//NewTransaction 
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
		return errInvalidUserID
	}

	if t.Description == "" {
		return errInvalidDescription
	}

	if t.Category == "" {
		return errInvalidCategory
	}

	if t.Date.IsZero() {
		return errInvalidDate
	}
	return nil
}

package usecase

import (
	"MoneyControl/internal/domain/entity"
	"MoneyControl/internal/domain/gateway"
	"context"
	"fmt"
	"time"
)

type TransactionUsecase struct {
	TransactionGateway gateway.TransactionGateway
}

type TransactionInputDTO struct {
	Value       float64   `json:"value"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Date        time.Time `json:"date"`
	UserID      uint      `json:"user_id"`
}

type TransactionOutputDTO struct {
	ID          uint      `json:"id"`
	Value       float64   `json:"value"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Date        time.Time `json:"date"`
	UserID      uint      `json:"user_id"`
}

func NewTransactionUsecase(transactionGateway gateway.TransactionGateway) *TransactionUsecase {
	return &TransactionUsecase{
		TransactionGateway: transactionGateway,
	}
}

func (t *TransactionUsecase) CreateTransaction(ctx context.Context, input TransactionInputDTO) (*TransactionOutputDTO, error) {
	transaction, err := entity.NewTransaction(input.Value, input.Description, input.Category, input.Date, input.UserID)
	if err != nil {
		return nil, fmt.Errorf("error creating new transaction: %v", err)
	}
	createdTransaction, err := t.TransactionGateway.CreateTransaction(transaction)
	if err != nil {
		return nil, fmt.Errorf("error inserting and creating new transaction: %v", err)
	}
	return &TransactionOutputDTO{
		ID:          createdTransaction.ID,
		Value:       createdTransaction.Value,
		Description: createdTransaction.Description,
		Category:    createdTransaction.Category,
		Date:        createdTransaction.Date,
		UserID:      createdTransaction.UserID,
	}, nil
}

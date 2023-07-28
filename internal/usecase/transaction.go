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

func toDTO(t *entity.Transaction) *TransactionOutputDTO {
	return &TransactionOutputDTO{
		ID:          t.ID,
		Value:       t.Value,
		Description: t.Description,
		Category:    t.Category,
		Date:        t.Date,
		UserID:      t.UserID,
	}
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
	output := toDTO(createdTransaction)
	return output, nil
}

func (t *TransactionUsecase) GetTransactionByID(id uint) (*TransactionOutputDTO, error) {
	transaction, err := t.TransactionGateway.ReadTransactionByID(id)
	if err != nil {
		return nil, fmt.Errorf("error reading transaction. Error: %v", err)
	}
	output := toDTO(transaction)
	return output, nil
}

// TODO: use pagination
func (t *TransactionUsecase) GetTransactionsPagination(limit, cursor int) ([]*TransactionOutputDTO, error) {
	transactionList, err := t.TransactionGateway.ReadTransactionsPagination(limit, cursor)
	if err != nil {
		return nil, fmt.Errorf("error retrieving all transactions: %v", err)
	}
	var transactionOutputList []*TransactionOutputDTO
	for _, transactionRes := range transactionList {
		transactionOutputList = append(transactionOutputList, toDTO(transactionRes))
	}
	return transactionOutputList, nil
}

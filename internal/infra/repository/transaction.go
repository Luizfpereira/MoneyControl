package repository

import (
	"context"

	"gorm.io/gorm"
	"MoneyControl/internal/domain/entity"
)

type TransactionRepositoryPSQL struct {
	Instance *gorm.DB
}

func NewTransactionRepositoryPSQL(instance *gorm.DB) *TransactionRepositoryPSQL {
	return &TransactionRepositoryPSQL{
		Instance: instance,
	}
}

func (r *TransactionRepositoryPSQL) CreateTransaction(ctx context.Context, t *entity.Transaction) error {
	return nil
}

package repository

import (
	"MoneyControl/internal/domain/entity"
	"fmt"

	"gorm.io/gorm"
)

type TransactionRepositoryPSQL struct {
	Instance *gorm.DB
}

func NewTransactionRepositoryPSQL(instance *gorm.DB) *TransactionRepositoryPSQL {
	return &TransactionRepositoryPSQL{
		Instance: instance,
	}
}

func (r *TransactionRepositoryPSQL) CreateTransaction(t *entity.Transaction) (*entity.Transaction, error) {
	record := r.Instance.Create(&t)
	fmt.Println(record)
	if record.Error != nil {
		return nil, record.Error
	}
	return t, nil
}

func (r *TransactionRepositoryPSQL) ReadTransactionByID(id uint) (*entity.Transaction, error) {
	return nil, nil
}

func (r *TransactionRepositoryPSQL) UpdateTransactionByID(id uint) (*entity.Transaction, error) {
	return nil, nil
}

func (r *TransactionRepositoryPSQL) DeleteTransactionByID(id uint) (uint, error) {
	return 0, nil
}

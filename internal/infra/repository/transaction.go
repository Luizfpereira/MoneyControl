package repository

import (
	"MoneyControl/internal/domain/entity"

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
	if record.Error != nil {
		return nil, record.Error
	}
	return t, nil
}

func (r *TransactionRepositoryPSQL) ReadTransactionByID(id uint) (*entity.Transaction, error) {
	var t *entity.Transaction
	result := r.Instance.Find(&t, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if err := t.Validate(); err != nil {
		return nil, err
	}
	return t, nil
}

func (r *TransactionRepositoryPSQL) UpdateTransactionByID(id uint, t *entity.Transaction) (*entity.Transaction, error) {
	r.Instance.Update()

	return nil, nil
}

func (r *TransactionRepositoryPSQL) DeleteTransactionByID(id uint) (uint, error) {
	return 0, nil
}

func (r *TransactionRepositoryPSQL) ReadTransactionsPagination(limit, cursos int) ([]*entity.Transaction, error) {
	return nil, nil
}

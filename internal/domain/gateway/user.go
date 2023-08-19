package gateway

type UserGateway interface {

}

CreateTransaction(transaction *entity.Transaction) (*entity.Transaction, error)
ReadTransactionByID(id uint) (*entity.Transaction, error)
// ReadTransactionsPagination(limit, cursos int) ([]*entity.Transaction, error)
UpdateTransactionByID(id uint) (*entity.Transaction, error)
DeleteTransactionByID(id uint) (uint, error)
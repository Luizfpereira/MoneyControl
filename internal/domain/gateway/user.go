package gateway

import "MoneyControl/internal/domain/entity"

type UserGateway interface {
	CreateUser(user *entity.User) (uint, error)
	ReadUserByID(id int) (*entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUserByID(id int) error
}

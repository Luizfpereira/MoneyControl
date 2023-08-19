package usecase

import (
	"MoneyControl/internal/domain/entity"
	"MoneyControl/internal/domain/gateway"
)

type UserUsecase struct {
	UserGateway gateway.UserGateway
}

func NewUserUsecase(gateway gateway.UserGateway) *UserUsecase {
	return &UserUsecase{UserGateway: gateway}
}

// type UserInputDTO struct {
// 	Name     string `json:"name"`
// 	LastName string `json:"last_name"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

type UserOutputDTO struct {
	ID       uint   `json:"ID"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
}

func toUserOutputDTO(u *entity.User, id uint) *UserOutputDTO {
	return &UserOutputDTO{
		ID:       id,
		Name:     u.Name,
		LastName: u.LastName,
		Email:    u.Email,
	}
}

func (u *UserUsecase) CreateUser(userInput *entity.User) (*UserOutputDTO, error) {
	res, err := u.UserGateway.CreateUser(userInput)
	if err != nil {
		return nil, err
	}
	return toUserOutputDTO(userInput, res), nil
}

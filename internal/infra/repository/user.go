package repository

import (
	"MoneyControl/internal/domain/entity"

	"gorm.io/gorm"
)

type UserRepositoryPSQL struct {
	Instance *gorm.DB
}

func NewUserRepositoryPSQL(instance *gorm.DB) *UserRepositoryPSQL {
	return &UserRepositoryPSQL{Instance: instance}
}

func (u *UserRepositoryPSQL) CreateUser(user *entity.User) (uint, error) {
	res := u.Instance.Create(&user)
	if res.Error != nil {
		return 0, res.Error
	}
	return user.ID, nil
}

func (u *UserRepositoryPSQL) ReadUserByID(id int) (*entity.User, error) {
	var user *entity.User
	res := u.Instance.Find(&user, id)
	if res.Error != nil {
		return nil, res.Error
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepositoryPSQL) UpdateUser(user *entity.User) error {
	res := u.Instance.UpdateColumns(user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (u *UserRepositoryPSQL) DeleteUserByID(id int) error {
	res := u.Instance.Delete(entity.User{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

package repository

import (
	"MoneyControl/internal/domain/entity"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserRepoSuite struct {
	suite.Suite
	repo *UserRepositoryPSQL
}

func TestUserRepoSuite(t *testing.T) {
	suite.Run(t, new(UserRepoSuite))
}

func (s *UserRepoSuite) SetupTest() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		s.Suite.T().Fatal(err)
	}
	db.AutoMigrate(entity.User{})

	repo := NewUserRepositoryPSQL(db)
	s.repo = repo

	u1, err := entity.NewUser("Luiz", "Pereira", "luiz.test@gmail.com", "*Mudar123")
	s.Assert().Nil(err)
	s.Assert().Nil(u1.Validate())

	createdUser1, err := repo.CreateUser(u1)
	s.Assert().Nil(err)
	s.Assert().Equal(u1, createdUser1)

	u2, err := entity.NewUser("Thais", "Montovani", "thais.montovani@gmail.com", "!Change456")
	s.Assert().Nil(err)
	s.Assert().Nil(u1.Validate())

	createdUser2, err := repo.CreateUser(u2)
	s.Assert().Nil(err)
	s.Assert().Equal(u2, createdUser2)
}

func (s *UserRepoSuite) TestCreateUser() {
	u3, err := entity.NewUser("Shirley", "Pereira", "shirley.test@gmail.com", "*Mudar123456")
	s.Assert().Nil(err)
	s.Assert().Nil(u3.Validate())
	res, err := s.repo.CreateUser(u3)
	s.Assert().Nil(err)
	s.Assert().Equal(u3, res)
}

func (s *UserRepoSuite) TestReadUserByID() {
	user1, err := s.repo.ReadUserByID(1)
	s.Assert().Nil(err)
	s.Assert().Equal(uint(1), user1.ID)
	s.Assert().Equal("Luiz", user1.Name)
	s.Assert().Equal("luiz.test@gmail.com", user1.Email)

	user2, err := s.repo.ReadUserByID(2)
	s.Assert().Nil(err)
	s.Assert().Equal(uint(2), user2.ID)
	s.Assert().Equal("Thais", user2.Name)
	s.Assert().Equal("thais.montovani@gmail.com", user2.Email)
}

func (s *UserRepoSuite) TestUpdateUserByID() {
	user1, err := s.repo.ReadUserByID(1)
	s.Assert().Nil(err)
	user1.Name = "Gabriel"
	user1.Password = "*Chama789"

	err = s.repo.UpdateUserByID(user1)
	s.Assert().Nil(err)
	updatedUser, err := s.repo.ReadUserByID(1)
	s.Assert().Nil(err)
	s.Assert().Equal(user1, updatedUser)
}

func (s *UserRepoSuite) TestDeleteUserByID() {
	err := s.repo.DeleteUserByID(1)
	s.Assert().Nil(err)

	user, err := s.repo.ReadUserByID(1)
	s.Assert().Nil(user)
	s.Assert().Error(err)
}

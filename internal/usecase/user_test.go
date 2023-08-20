package usecase

import (
	"MoneyControl/internal/domain/entity"
	"MoneyControl/internal/domain/gateway/mocks"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type UserUsecaseSuite struct {
	suite.Suite
	mockRepo *mocks.MockUserGateway
}

func TestUserRepoSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseSuite))
}

func (s *UserUsecaseSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	mockUserRepo := mocks.NewMockUserGateway(ctrl)
	s.mockRepo = mockUserRepo
}

func (s *UserUsecaseSuite) TestCreateUser() {
	user := entity.User{
		Name:     "Luiz",
		LastName: "Pereira",
		Email:    "luiz.test@gmail.com",
		Password: "*Mudar123",
	}
	s.mockRepo.EXPECT().CreateUser(&user).Return(uint(1), nil)
	userOutputDTO := toUserOutputDTO(&user, 1)

	userUsecase := NewUserUsecase(s.mockRepo)
	userOutput, err := userUsecase.CreateUser(&user)
	s.Assert().Nil(err)
	s.Assert().Equal(userOutputDTO, userOutput)
}

func (s *UserUsecaseSuite) TestReadUserByID() {
	user := entity.User{
		ID:       uint(1),
		Name:     "Luiz",
		LastName: "Pereira",
		Email:    "luiz.test@gmail.com",
		Password: "*Mudar123",
	}
	s.mockRepo.EXPECT().ReadUserByID(1).Return(&user, nil)
	userOutputDTO := toUserOutputDTO(&user, uint(1))

	userUsecase := NewUserUsecase(s.mockRepo)
	userOutput, err := userUsecase.ReadUserByID(1)
	s.Assert().Nil(err)
	s.Assert().Equal(userOutputDTO, userOutput)
}

func (s *UserUsecaseSuite) TestUpdateUser() {
	user := entity.User{
		ID:       uint(1),
		Name:     "Luiz",
		LastName: "Pereira",
		Email:    "luiz.test@gmail.com",
		Password: "*Mudar123",
	}
	s.mockRepo.EXPECT().ReadUserByID(1).Return(&user, nil)
	toUpdate := entity.User{
		ID:   user.ID,
		Name: "Gabriel",
	}

	s.mockRepo.EXPECT().UpdateUser(&toUpdate).Return(nil)

	userUsecase := NewUserUsecase(s.mockRepo)
	err := userUsecase.UpdateUser(&toUpdate)
	s.Assert().Nil(err)
}

func (s *UserUsecaseSuite) TestDeteleUserByID() {
	s.T().Run("testing success deleting ID 1", func(t *testing.T) {
		s.mockRepo.EXPECT().DeleteUserByID(1).Return(nil)
		userUsecase := NewUserUsecase(s.mockRepo)
		err := userUsecase.DeleteUserByID(1)
		s.Assert().Nil(err)
	})

	s.T().Run("testing failure in deleting ID 2", func(t *testing.T) {
		s.mockRepo.EXPECT().DeleteUserByID(2).Return(errors.New("error deleting user id 2"))
		userUsecase := NewUserUsecase(s.mockRepo)
		err := userUsecase.DeleteUserByID(2)
		s.Assert().NotNil(err)
	})

}

package usecase

import (
	"MoneyControl/internal/domain/entity"
	"MoneyControl/internal/domain/gateway/mocks"
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

package web

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserHandlerSuite struct {
	suite.Suite
}

func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerSuite))
}

//https://www.cloudbees.com/blog/testing-http-handlers-go - ler para testar

func (u *UserHandlerSuite) TestGetUserByID() {
	r, err := http.NewRequest("GET", "/users/1", nil)
}

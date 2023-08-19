package repository

import (
	"MoneyControl/internal/domain/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//para testar, configurar banco de dados de teste do tipo sqlite (em memoria)

type TransactionRepoSuite struct {
	db *gorm.DB
	suite.Suite
}

func TestTransactionRepoSuite(t *testing.T) {
	suite.Run(t, new(TransactionRepoSuite))
}

func (s *TransactionRepoSuite) SetupTest() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		s.Suite.T().Fatal(err)
	}
	db.AutoMigrate(entity.Transaction{})
	s.db = db

	repo := NewTransactionRepositoryPSQL(s.db)
	t, err := entity.NewTransaction(40.6, "gasto", "mercado", time.Now(), 1)
	if err != nil {
		s.Suite.T().Fatal(err)
	}
	result, err := repo.CreateTransaction(t)
	s.Assert().Nil(err)
	s.Assert().Equal(result, t)

	t2, err := entity.NewTransaction(10, "compras", "shopping", time.Now(), 2)
	if err != nil {
		s.Suite.T().Fatal(err)
	}
	result2, err := repo.CreateTransaction(t2)
	s.Assert().Nil(err)
	s.Assert().Equal(result2, t2)
}

func (s *TransactionRepoSuite) TestCreateTransaction() {
	repo := NewTransactionRepositoryPSQL(s.db)
	t, err := entity.NewTransaction(20.5, "gasto 2", "mercado", time.Now(), 2)
	if err != nil {
		s.Suite.T().Fatal(err)
	}
	result, err := repo.CreateTransaction(t)
	s.Assert().Nil(err)
	s.Assert().Equal(result, t)
}

func (s *TransactionRepoSuite) TestReadTransactionByID() {
	repo := NewTransactionRepositoryPSQL(s.db)
	t, err := repo.ReadTransactionByID(1)
	s.Assert().Nil(err)
	s.Assert().Equal(t.ID, uint(1))
	s.Assert().Equal(40.6, t.Value)
}

func (s *TransactionRepoSuite) TestUpdateTransactionByIDValue() {
	repo := NewTransactionRepositoryPSQL(s.db)
	t, err := repo.ReadTransactionByID(1)
	s.Assert().Nil(err)

	s.Suite.T().Run("Testing value update", func(ts *testing.T) {
		t.Value = 5
		err := repo.UpdateTransactionByID(1, t)
		s.Assert().Nil(err)
		res, err := repo.ReadTransactionByID(1)
		s.Assert().Nil(err)
		s.Assert().Equal(t, res)
	})

	s.Suite.T().Run("Testing value, description, category, date, userID update", func(ts *testing.T) {
		t.Value = 550.6
		t.Category = "categoria update"
		t.Date = time.Now().UTC()
		t.UserID = 3
		err := repo.UpdateTransactionByID(1, t)
		s.Assert().Nil(err)
		res, err := repo.ReadTransactionByID(1)
		s.Assert().Nil(err)
		s.Assert().Equal(t, res)
	})
}

func (s *TransactionRepoSuite) TestDeleteTransactionByID() {
	repo := NewTransactionRepositoryPSQL(s.db)
	err := repo.DeleteTransactionByID(1)
	s.Assert().Nil(err)

	t, err := repo.ReadTransactionByID(1)
	s.Assert().Nil(t)
	s.Assert().Error(err)
}

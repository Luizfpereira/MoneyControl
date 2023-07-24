package main

import (
	"MoneyControl/internal/infra/database"
	"MoneyControl/internal/infra/repository"
	"MoneyControl/internal/usecase"
	"context"
	"fmt"
	"time"
)

func main() {
	instance := database.ConnectSingleton()
	database.Migrate(instance)

	transactionRep := repository.NewTransactionRepositoryPSQL(instance)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRep)
	// t, err := entity.NewTransaction(12.6, "teste", "farmacia", time.Now().UTC(), 1)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	input := usecase.TransactionInputDTO{12.6, "teste", "farmacia", time.Now().UTC(), 1}
	output, err := transactionUsecase.CreateTransaction(context.Background(), input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output)
}

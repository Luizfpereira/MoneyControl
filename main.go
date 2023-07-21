package main

import (
	"MoneyControl/internal/infra/database"
)

func main() {
	instance := database.ConnectSingleton()
	database.Migrate(instance)
}

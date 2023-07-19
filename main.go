package main

import (
	"MoneyControl/infra/database"
	"fmt"
)

func main() {
	instance := database.ConnectSingleton()
	database.Migrate(instance)
	instance.Select("Select * from users;")
	fmt.Println("teste ", instance)
}

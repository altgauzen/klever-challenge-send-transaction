package main

import (
	"fmt"
	"klever/api/database"
	"klever/api/models"
	"klever/api/routes"
)


func main() {
	models.Details =[]models.Detail {
		{Address: "bc1qyzxdu4px4jy8gwhcj82zpv7qzhvc0fvumgnh0r", Balance: "144011754", TotalTx: 17000},
		{Address: "de1qyzxdu4px4jy8gwhcj82zpv7qzhvc0fvumgnh0r", Balance: "288011754", TotalTx: 18000},
	}
	models.Balances =[]models.Balance {
		{Confirmed: "1312321321321", Unconfirmed: "12321321"},
		{Confirmed: "1412321321321", Unconfirmed: "12345621"},
	}
	database.ConectaComBancoDeDados()

	fmt.Println("Starting server...")
	routes.HandleRequest()
}

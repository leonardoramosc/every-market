package main

import (
	"github.com/leonardoramosc/every-market/internal/database"
	"github.com/leonardoramosc/every-market/internal/routes"
)

func main() {
	database.Connect()
	database.RunMigrations()

	routes.InitRouter()
}

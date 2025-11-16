package main

import (
	"log"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/entities/authentication"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/config"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/database"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/entities/people"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/server" // Use your module name
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := database.InitDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to load database")
	}

	// start the server
	router := server.SetupRouter()

	// Define the routes
	authentication.RegisterRoutes(router, db)
	people.RegisterRoutes(router, db)

	router.Run(":8080")
}

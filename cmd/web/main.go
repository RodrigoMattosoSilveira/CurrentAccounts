package main

import (
	"log"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/config"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/database"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/cmd/web/fiberapp"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/cmd/web/ginapp"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/cmd/web/proxy"
)	

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := database.InitDatabase(config.Cfg.DB_NAME); 
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	// Start Gin server
	go ginapp.StartGin(config.Cfg.GIN_PORT, db)

	// Start Fiber server
	go fiberapp.StartFiber(config.Cfg.FIBER_PORT, db)

	// Start reverse proxy (entrypoint)
	proxy.StartProxy(
		config.Cfg.PROXY_PORT,
		config.Cfg.GIN_PORT,
		config.Cfg.FIBER_PORT,
	)
}

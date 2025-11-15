package main

import (
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/routes"
  "github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/server"  // Use your module name
)

func main() {
  r := server.SetupRouter()
  routes.SetupRoutes(r)
  r.Run(":8080")
}

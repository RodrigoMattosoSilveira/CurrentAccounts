package main

import (
	"github.com/RodrigoMattosoSilveira/mygo/internal/routes"
  "github.com/RodrigoMattosoSilveira/mygo/internal/server"  // Use your module name
)

func main() {
  r := server.SetupRouter()
  routes.SetupRoutes(r)
  r.Run(":8080")
}

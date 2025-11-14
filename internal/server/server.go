package server

import (
	"log"
	"path/filepath"

	"github.com/RodrigoMattosoSilveira/mygo/internal/utilities"
	"github.com/gin-gonic/gin"
)


func SetupRouter()  *gin.Engine {
  router :=gin.Default()

    // Determine project root (2 levels up from cmd/web)
    projectRoot, err := utilities.FindProjectRoot()
    if err != nil {
        log.Fatalf("failed to resolve project root: %v", err)
    }

    // Static files (e.g. CSS, JS, images)
    staticPath := filepath.Join(projectRoot, "static")
    router.Static("/static", staticPath)

  return router
}

package server

import (
	"log"
	"path/filepath"

  "github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/utilities"
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

  // Sessions (cookie store)
  store := cookie.NewStore([]byte("very-secret-session-key"))
  store.Options(sessions.Options{
    Path:     "/",
    HttpOnly: true,
    // Secure: true, // enable in production with HTTPS
  })
  router.Use(sessions.Sessions("app_session", store))

  return router
}

package templates

import (
	"log"
	"path/filepath"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/utilities" // IMPORTANT: Use your actual go.mod module name

	"github.com/gin-gonic/gin"
)

// LoadTemplates finds and loads all HTML templates from the nested 'templates'
// directory into the Gin engine. It uses a recursive glob pattern.
func Load(engine *gin.Engine) {
	// 1. Find the project root directory.
	projectRoot, err := utilities.FindProjectRoot()
	if err != nil {
		log.Fatalf("FATAL: Could not find project root: %v", err)
	}

	// 2. Construct the full path to the templates directory and the glob pattern.
	// The pattern `**/*.tmpl` will match all files ending in .tmpl within the
	// 'templates' directory and all its subdirectories.
	templatePattern := filepath.Join(projectRoot, "templates", "**", "*.tmpl")
	log.Printf("INFO: Loading templates from pattern: %s", templatePattern)

	// 3. Load the templates. Gin will use the file's base name as the template name.
	// For example, 'templates/root/home/welcome.tmpl' will be named 'welcome.tmpl'.
	engine.LoadHTMLGlob(templatePattern)
}

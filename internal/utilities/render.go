package utilities

import (
	"html/template"
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// RenderTemplate dynamically parses and executes a set of templates.
// It now correctly assumes templates are located in 'internal/templates'.
func RenderTemplate(c *gin.Context, name string, data gin.H, files ...string) {
	// 1. Find the project root.
	projectRoot, err := FindProjectRoot()
	if err != nil {
		log.Printf("ERROR: Failed to find project root: %v", err)
		c.AbortWithStatus(500)
		return
	}

	// 2. Create a slice of absolute paths for all requested template files.
	absFiles := make([]string, len(files))
	for i, file := range files {
		// THE FIX IS HERE: We add "internal" to the path construction.
		absFiles[i] = filepath.Join(projectRoot, "internal", "templates", file)
	}

	// 3. Parse the template files.
	tmpl, err := template.New(name).ParseFiles(absFiles...)
	if err != nil {
		log.Printf("ERROR: Failed to parse templates %v: %v", absFiles, err)
		c.AbortWithStatus(500)
		return
	}

	// 4. Execute the template.
	err = tmpl.Execute(c.Writer, data)
	if err != nil {
		log.Printf("ERROR: Failed to execute template '%s': %v", name, err)
		c.AbortWithStatus(500)
	}
}

func RenderModalDialog(c *gin.Context, title string, body string) {
	data := gin.H{
		"title":        title,
		"body":         body,
		"action_route": "", //
		"action_label": "",
		"action_class": "",
	}
	// Trigger a dialog_event in the server!
	c.Header("HX-Retarget", "#htmx-server-dialog-container")
	c.Header("HX-Reswap", "innerHTML")
	c.Header("HX-Trigger", "dialog_event")
	templateFiles := []string{
		"root/general/modalDialog.tmpl",
	}
	RenderTemplate(c, "modalDialog", data, templateFiles...)
}

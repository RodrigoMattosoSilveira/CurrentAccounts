package utilities

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/gin-gonic/gin"

	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/constants"
)

// Helper: renders a view with layout and partials
// templateCache avoids re-parsing templates repeatedly
var templateCache = struct {
	mu   sync.RWMutex
	data map[string]*template.Template
}{
	data: make(map[string]*template.Template),
}

var TMPL_FOLDER = constants.TMPL_FOLDER

func Render(c *gin.Context, partial string, data gin.H) {
	var layout string
	var layoutName string
	route := c.FullPath()
	wd, _ := os.Getwd()
	templDir := filepath.Join(wd, TMPL_FOLDER)

	switch route {
	case "/":
		layout = "layout.tmpl"
		layoutName = "layout"
	case "/about":
		layout = "layout.tmpl"
		layoutName = "layout"
	case "/welcome":
		layout = "body.tmpl"
		layoutName = "body"
	case "/bemvindo":
		layout = "body.tmpl"
		layoutName = "body"
	case "/login":
		layout = "body.tmpl"
		layoutName = "body"
	case "/logon":
		layout = "body.tmpl"
		layoutName = "body"
	default:
		layout = "layout.tmpl"
		layoutName = "layout"
	}
	// Key for cache
	key := layout + "|" + partial

	// Try cached template
	templateCache.mu.RLock()
	t, ok := templateCache.data[key]
	templateCache.mu.RUnlock()

	if !ok {
		files := []string{
			filepath.Join(templDir, layout),
			filepath.Join(templDir, partial),
		}
		t = template.Must(template.ParseFiles(files...))
		templateCache.mu.Lock()
		templateCache.data[key] = t
		templateCache.mu.Unlock()
	}

	// Execute template using its defined name (not filename)
	c.Status(http.StatusOK)
	if err := t.ExecuteTemplate(c.Writer, layoutName, data); err != nil {
		c.String(http.StatusInternalServerError, "template error: %v", err)
	}
}

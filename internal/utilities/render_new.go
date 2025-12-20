package utilities

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/utils"
		"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/config"
)

// Helper: renders a view with layout and partials
// templateCache avoids re-parsing templates repeatedly
var templateCache = struct {
	mu   sync.RWMutex
	data map[string]*template.Template
}{
	data: make(map[string]*template.Template),
}

type TmplPartial struct {
	Fn       string
	Name     string
	FullName string
}

func RenderPage(c *fiber.Ctx, partials []TemplateStruct, data fiber.Map) error {
	tmplRoot := config.Cfg.TMPL_ROOT
	projectRoot, err := utils.FindProjectRoot()
	if err != nil {
		log.Printf("ERROR: Failed to find project root: %v", err)		
		// c.AbortWithStatus(500)
		return err
	}

	// 3. Parse the templates.
	var partialsStr []string
	partialsStrKey := ""
	for _, partial := range partials {
		partial.FullName = filepath.Join(projectRoot, tmplRoot, partial.Fn)
		partialsStr = append(partialsStr, readTemplateFile(partial))
		partialsStrKey += partial.FullName
	}

	// Check cache first
	if cachedTmpl, exists := getCachedTemplate(partialsStrKey); exists {
		// Execute cached template into the response body
		c.Type("html", "utf-8")
		return cachedTmpl.Execute(c.Response().BodyWriter(), data)
	}

	// Not in cache, parse anew and cache it
	var tmpl *template.Template			
	tmpl = template.New("layout")
	for _, part := range partialsStr {
		tmpl, err = tmpl.Parse(part)
		if err != nil {
			log.Fatal(err)
		}
	}
	// Cache the parsed template
	cacheTemplate(partialsStrKey, tmpl)

	// Execute template into the response body
	// Use c.Response().BodyWriter() so template writes directly to Fiber response
	c.Type("html", "utf-8")
	return tmpl.Execute(c.Response().BodyWriter(), data)
}
func RenderPartial(c *fiber.Ctx, templateFn string, data fiber.Map) error {
	partial := TemplateStruct{
		Fn:   templateFn,
		Name: "partial",
		FullName: "",
	}		

	tmplRoot := config.Cfg.TMPL_ROOT
	projectRoot, err := utils.FindProjectRoot()
	if err != nil {
		log.Printf("ERROR: Failed to find project root: %v", err)		
		// c.AbortWithStatus(500)
		return err
	}

	partial.FullName = filepath.Join(projectRoot, tmplRoot, partial.Fn)
	tmplString := readTemplateFile(partial)	

	tmpl := template.New(partial.Name)
	tmpl, err = tmpl.Parse(tmplString)
	if err != nil {
		log.Fatal(err)
	}

	c.Type("html", "utf-8")
	return tmpl.Execute(c.Response().BodyWriter(), data)
}
func readTemplateFile(tmpl TemplateStruct) string {
	// Read the file into a byte slice, then convert to string
	content, err := os.ReadFile(tmpl.FullName)
	if err != nil {
		log.Fatal(err)
	}
	// templateStr := tmpl.Prefix + string(content) + "\n" + "{{ end }}"
	// templateStr := "\n" + tmpl.Prefix + "\n" + string(content)+ "\n"  + "{{ end }}"
	templateStr := derivePrefix(tmpl.Name) + string(content) + "\n" + "{{ end }}"
	// log.Println(templateStr)
	return templateStr
}

func derivePrefix(name string) string {
	// Prefix: `{{ define "bottom" }}`
	return "\n" + `{{ define "` + name + `" }}` + "\n"
}

func getCachedTemplate(key string) (*template.Template, bool) {
	templateCache.mu.RLock()
	tmpl, exists := templateCache.data[key]
	templateCache.mu.RUnlock()
	
	return tmpl, exists
}


func cacheTemplate(key string, tmpl *template.Template) {
	templateCache.mu.Lock()
	templateCache.data[key] = tmpl
	templateCache.mu.Unlock()
}

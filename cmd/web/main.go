package main

import (
  "bytes"
  "html/template"
  "net/http"
  "github.com/gin-gonic/gin"
)
// Template data structure
type PageData struct {
  Title string
  Body  string
}

func setupRouter() *gin.Engine {
	router := gin.Default()
  router.SetTrustedProxies([]string{"192.168.1.2"})

	// Load templates
	router.SetHTMLTemplate(template.Must(template.ParseFiles("../../internal/templates/hello.tmpl")))
  
  router.GET("/hello", HelloHandler)

	return  router
}

// RenderTemplate renders the HTML template with provided data
func RenderTemplate(tmplStr string, data PageData) (string, error) {
  tmpl, err := template.New("page").Parse(tmplStr)
  if err != nil {
    return "", err
  }

  var buf bytes.Buffer
  if err := tmpl.Execute(&buf, data); err != nil {
    return "", err
  }

  return buf.String(), nil
}

func main() {

  r := setupRouter()
	r.Run(":8080")
}

func HelloHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "hello.tmpl", gin.H{
		"Title": "Hello, Gin!",
    "Body":  "Welcome to the Gin web framework.",
	})
}
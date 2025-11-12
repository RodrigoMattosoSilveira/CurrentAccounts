package main

import (
  "bytes"
  "html/template"

  "github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/routes"
   "github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/server"
)
// Template data structure
type PageData struct {
  Title string
  Body  string
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
  r := server.SetupRouter()
  routes.SetupRoutes(r)
  r.Run(":8080")
}


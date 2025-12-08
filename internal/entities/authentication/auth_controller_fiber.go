package authentication

import (
	"github.com/gofiber/fiber/v2"

	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/utilities"
)


func (app *PeopleController) ShowFiber(c *fiber.Ctx) error {
	templateFiles := []string{
		"root/layout.tmpl",
		"root/authentication/login.tmpl",
	}

	// Call our custom renderer.
	// The name "layout.tmpl" tells the template engine which template definition to execute first.
	utilities.RenderTemplateFiber(c, "layout", map[string]any{
		"Tenant": "MC",
		"Host":   "Madone Logistics",
	}, templateFiles...)
	return nil
}
func (app *PeopleController) ShowLogin(c *fiber.Ctx) error {
	templateFiles := []string{
		"root/layout.tmpl",
		"root/authentication/login.tmpl",
	}

	// Call our custom renderer.
	// The name "layout.tmpl" tells the template engine which template definition to execute first.
	utilities.RenderTemplateFiber(c, "layout", map[string]any{
		"Tenant": "MC",
		"Host":   "Madone Logistics",
	}, templateFiles...)
	return nil
}
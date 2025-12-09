package authentication

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	k "github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/constants"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/entities/people"
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

func (app *PeopleController) HandleLogin(c *fiber.Ctx) error {
	type formS struct {
		Email string
		Password string
	}
	var form formS
	templateFiles := []string{
		"root/layout.tmpl",
		"root/authentication/welcome.tmpl",
	}

	// var loginForm LoginForm
	// if err := c.ShouldBind(&loginForm); err != nil {
	// 	utilities.RenderModalDialog(c, "Invalid login form", "Please try again")
	// }
	if err := c.BodyParser(&form); err != nil {
		
	 }
	var person people.Person
	person, err := app.service.GetByEmail(form.Email)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		utilities.RenderTemplateFiber(c, "layout", map[string]any{
			"Tenant": "MC",
			"Host":   "Madone Logistics",
			"Error": "User not found",	
			"Message": "Try again",
		}, templateFiles...)
		return nil
	}

	if !checkPasswordHash(person.Password, form.Password) {
		c.Status(http.StatusUnauthorized)
		utilities.RenderTemplateFiber(c, "layout", map[string]any{
			"Tenant": "MC",
			"Host":   "Madone Logistics",
			"Error": "Invalid password",	
			"Message": "Try again",
		}, templateFiles...)
		return nil
	}

	sess := utilities.GetSession(c)
	sess.Set(k.PERSON_ID, person.ID)

	if err := sess.Save(); err != nil {
		c.Status(http.StatusInternalServerError)
		utilities.RenderTemplateFiber(c, "layout", map[string]any{
			"Tenant": "MC",
			"Host":   "Madone Logistics",
			"Error": "Failed to save session",	
		}, templateFiles...)
		return nil
	}

	// This forces HTMX to reload the whole page without treating it as a fragment
	c.Status(http.StatusOK)
	utilities.RenderTemplateFiber(c, "layout", map[string]any{
		"Tenant": "MC",
		"Host":   "Madone Logistics",
	}, templateFiles...)
	return nil

}

func (app *PeopleController) HandleLogout(c *fiber.Ctx) error {
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

package authentication

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"

	k "github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/constants"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/entities/people"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/utilities"
)

type Map map[string]any
type RedirectConfig struct {
	Params  Map               // Route parameters
	Queries map[string]string // Query map
}

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
	type formStruct struct {
		Email string
		Password string
	}
	var form formStruct

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
		templateFiles := []string{"root/layout.tmpl", "root/authentication/login.tmpl", }
		return utilities.RenderTemplateFiber(c, "layout", map[string]any{
			"Tenant": "MC", "Host":   "Madone Logistics", "Error": "User not found, try again",
		}, templateFiles...)	
	}

	if !checkPasswordHash(person.Password, form.Password) {
		c.Status(http.StatusUnauthorized)
		templateFiles := []string{"root/layout.tmpl", "root/authentication/login.tmpl", }
		return utilities.RenderTemplateFiber(c, "layout", map[string]any{
			"Tenant": "MC", "Host":   "Madone Logistics", "Error": "Invalid password, try again",
		}, templateFiles...)	
	}

	sess := utilities.GetSession(c)
	sess.Set(k.PERSON_ID, person.ID)

	if err := sess.Save(); err != nil {
		c.RedirectToRoute("login", map[string]any {"Error": "Internal Server Error"}, http.StatusInternalServerError)
		return nil
	}

	// This forces HTMX to reload the whole page without treating it as a fragment
	c.Status(http.StatusOK)
	// queries := map[string]string {"email": person.Email}
	// c.RedirectToRoute("HandleWelcome", map[string]any {"queries": queries}, http.StatusOK)
	// return c.Redirect("/welcome?Tenant=MC?Host=Madone%20Logistics?email=miguel_moraes@camilapassos.com.br", http.StatusOK)
	templateFiles := []string{
		"root/layout.tmpl",
		"root/authentication/welcome.tmpl",
	}

	return utilities.RenderTemplateFiber(c, "layout", map[string]any{
		"Tenant": "MC",
		"Host":   "Madone Logistics",
		"Name": person.Name,
	}, templateFiles...)
}

func (app *PeopleController) HandleLogout(c *fiber.Ctx) error {
	c.RedirectToRoute("login", map[string]any {}, http.StatusOK)

	return nil
}
func (app *PeopleController) HandleWelcome(c *fiber.Ctx) error{
	email := c.Query("email")
	var person people.Person

	templateFiles := []string{
		"root/layout.tmpl",
		"root/authentication/welcome.tmpl",
	}
	person, err := app.service.GetByEmail(email)
	if err != nil {
		log.Println("Error fetching person by email:", err)
		c.Status(http.StatusUnauthorized)	
		return utilities.RenderTemplateFiber(c, "layout", map[string]any{
			"Tenant": "MC",
			"Host":   "Madone Logistics",
			"Error": "Invlid password",	
			"Message": "Try again",
		}, templateFiles...)
	}
	c.Status(http.StatusOK)
	return utilities.RenderTemplateFiber(c, "layout", map[string]any{
		"Tenant": "MC",
		"Host":   "Madone Logistics",
		"Name": person.Name,
	}, templateFiles...)
}
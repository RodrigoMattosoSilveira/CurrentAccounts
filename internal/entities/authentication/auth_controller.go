package authentication

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	k "github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/constants"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/entities/people"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/utilities"
)
type PeopleController struct {
	service people.Service
}

type LoginForm struct {
	Email    string
	Password string
}
const currentUserKey = "currentUser"
type Map map[string]any
type RedirectConfig struct {
	Params  Map               // Route parameters
	Queries map[string]string // Query map
}

func NewPeopleController(service people.Service) *PeopleController {
	// forcing a change to test git
	return &PeopleController{
		service: service,
	}
}

func NewController(db *gorm.DB) *AppFiber {
	return &AppFiber{DB: db}
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
	if err := c.BodyParser(&form); err != nil {
		c.Status(http.StatusInternalServerError)
		templateFiles := []string{"root/layout.tmpl", "root/authentication/login.tmpl", }
		return utilities.RenderTemplateFiber(c, "layout", map[string]any{
			"Tenant": "MC", "Host":   "Madone Logistics", "Error": "Unable to parse form, try again; inform the administrator if the problem persists",
		}, templateFiles...)	
	}

	var person people.Person
	person, err := app.service.GetByEmail(form.Email)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity)
		templateFiles := []string{"root/layout.tmpl", "root/authentication/login.tmpl", }
		return utilities.RenderTemplateFiber(c, "layout", map[string]any{
			"Tenant": "MC", "Host":   "Madone Logistics", "Error": "User not found, try again",
		}, templateFiles...)	
	}

	if !checkPasswordHash(person.Password, form.Password) {
		// see here for a good discussion on status codes for invalid login
		// https://stackoverflow.com/questions/7939137/what-http-status-code-should-be-used-for-wrong-input
		c.Status(http.StatusUnprocessableEntity)
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

func (app *PeopleController) ShowLogon(c *fiber.Ctx) error {
	// We need the layout and the specific welcome page.
	// The paths are relative to the 'templates' directory.
	templateFiles := []string{
		"root/layout.tmpl",
		"root/authentication/logon.tmpl",
	}

	// Call our custom renderer.
	// The name "layout.tmpl" tells the template engine which template definition to execute first.
	return utilities.RenderTemplateFiber(c, "layout", map[string]any{
		"Tenant": "MC",
		"Host":   "Madone Logistics",
	}, templateFiles...)
}

func (app *PeopleController) HandleRegister(c *fiber.Ctx) error {
	var person people.Person

	type formStruct struct {
		Fullname string
		Address string
		Cell string
		Email string
		Password string
	}
	var form formStruct
	if err := c.BodyParser(&form); err != nil {
		c.Status(http.StatusInternalServerError)
		templateFiles := []string{"root/layout.tmpl", "root/authentication/logon.tmpl", }
		return utilities.RenderTemplateFiber(c, "layout", map[string]any{
			"Tenant": "MC", "Host":   "Madone Logistics", "Error": "Unable to parse form, try again; inform the administrator if the problem persists",
		}, templateFiles...)	
	}

	// TODO add validate.validator
	name := form.Fullname
	address := form.Address
    email := form.Email
	cell := form.Cell
    password := form.Password

	if name == "" {
		c.Status(http.StatusUnprocessableEntity)
		templateFiles := []string{"root/layout.tmpl", "root/authentication/logon.tmpl", }
		return utilities.RenderTemplateFiber(c, "layout", map[string]any{
			"Tenant": "MC", "Host":   "Madone Logistics", "Error": "Invalid name, try again",
		}, templateFiles...)	

	}

	if address == "" {
		c.Status(http.StatusUnprocessableEntity)
		templateFiles := []string{"root/layout.tmpl", "root/authentication/logon.tmpl", }
		return utilities.RenderTemplateFiber(c, "layout", map[string]any{
			"Tenant": "MC", "Host":   "Madone Logistics", "Error": "Invalid address, try again",
		}, templateFiles...)	

	}

	if email == "" {
		c.Status(http.StatusUnprocessableEntity)
		templateFiles := []string{"root/layout.tmpl", "root/authentication/logon.tmpl", }
		return utilities.RenderTemplateFiber(c, "layout", map[string]any{
			"Tenant": "MC", "Host":   "Madone Logistics", "Error": "Invalid email, try again",
		}, templateFiles...)	

	}

	if password == "" {
		c.Status(http.StatusUnprocessableEntity)
		templateFiles := []string{"root/layout.tmpl", "root/authentication/logon.tmpl", }
		return utilities.RenderTemplateFiber(c, "layout", map[string]any{
			"Tenant": "MC", "Host":   "Madone Logistics", "Error": "Invalid password, try again",
		}, templateFiles...)	
	}

	person, err := app.service.GetByEmail(email)
	if err == nil {
		// see here for a good discussion on status codes for invalid login
		// https://stackoverflow.com/questions/7939137/what-http-status-code-should-be-used-for-wrong-input
		c.Status(http.StatusConflict)
		templateFiles := []string{"root/layout.tmpl", "root/authentication/logon.tmpl", }
		return utilities.RenderTemplateFiber(c, "layout", map[string]any{
			"Tenant": "MC", "Host":   "Madone Logistics", "Error": "Person record already registered, try again",
		}, templateFiles...)	
	}
 
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
		c.Status(http.StatusInternalServerError)
		templateFiles := []string{"root/log.tmpl", "root/authentication/logon.tmpl", }
		return utilities.RenderTemplateFiber(c, "layout", map[string]any{
			"Tenant": "MC", "Host":   "Madone Logistics", "Error": "Invalid password, try again",
		}, templateFiles...)	
    }

    person = people.Person{
		Name: name,
		Address: address,
        Email:email,
		Cell: cell,
        Password: string(hash),
        Role: "Person",
    }
	err = app.service.Create(&person)
    if err != nil {
		c.Status(http.StatusInternalServerError)
		templateFiles := []string{"root/log.tmpl", "root/authentication/logon.tmpl", }
		message := "Unable to create person record, " + err.Error()
		return utilities.RenderTemplateFiber(c, "layout", map[string]any{
			"Tenant": "MC", "Host":   "Madone Logistics", "Error": message,
		}, templateFiles...)	
    }

	c.Status(http.StatusOK)
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
func (app *PeopleController) HandleNewPwd(c *fiber.Ctx) error {
	return nil
}


// CurrentPerson retrieves the logged-in person from session (or nil).
func (app *PeopleController) CurrentPerson(c *fiber.Ctx) *people.Person {
	// val := c.Get(currentUserKey); 
	// if val != "" {
	// 	if u, ok := val.(*people.Person); ok {
	// 		return u
	// 	}
	// }

	sess := utilities.GetSession(c)
	idVal := sess.Get(k.PERSON_ID)
	if idVal == nil {
		return nil
	}

	var idUint uint
	switch v := idVal.(type) {
	case uint:
		idUint = v
	case int:
		if v > 0 {
			idUint = uint(v)
		}
	default:
		return nil
	}

	var person people.Person
	person, err := app.service.GetByID(idUint)
	if err != nil {
		return nil
	}

	return &person
}

func checkPasswordHash(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

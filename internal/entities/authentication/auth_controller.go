package authentication

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/constants"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/entities/people"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/utilities"
)
const currentUserKey = "currentUser"
type LoginForm struct {
	Email    string
	Password string
}

type PeopleController struct {
	service people.Service
}

func NewPeopleController(service people.Service) *PeopleController {
	// forcing a change to test git
	return &PeopleController{
		service: service,
	}
}

func NewController(db *gorm.DB) *App {
	return &App{DB: db}
}

func (app *PeopleController) ShowLogin(c *gin.Context) {
	templateFiles := []string{
		"root/layout.tmpl",
		"root/authentication/login.tmpl",
	}

	// Call our custom renderer.
	// The name "layout.tmpl" tells the template engine which template definition to execute first.
	utilities.RenderTemplate(c, "layout", gin.H{
		"Tenant": "MC",
		"Host":   "Madone Logistics",
	}, templateFiles...)
}
func (app *PeopleController) HandleLogin(c *gin.Context) {

	// var loginForm LoginForm
	// if err := c.ShouldBind(&loginForm); err != nil {
	// 	utilities.RenderModalDialog(c, "Invalid login form", "Please try again")
	// }
	email := c.PostForm("email")
	password := c.PostForm("password")

	var person people.Person
	person, err := app.service.GetByEmail(email)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		utilities.RenderModalDialog(c, "Invalid email", "Please try again")
		return
	}

	if !CheckPasswordHash(person.Password, password) {
		c.Status(http.StatusUnauthorized)
		utilities.RenderModalDialog(c, "Invalid password", "Please try again")
		return
	}

	sess := sessions.Default(c)
	sess.Set(constants.PERSON_ID, person.ID)
	if err := sess.Save(); err != nil {
		utilities.RenderModalDialog(c, "Failed to save session", "Please try again")
		return
	}

	// This forces HTMX to reload the whole page without treating it as a fragment
	c.Status(http.StatusOK)
	c.Header("HX-Redirect", "/welcome/?email="+ person.Email)
}
func (app *PeopleController) HandleWelcome(c *gin.Context) {
	email := c.Query("email")
	var person people.Person
	person, err := app.service.GetByEmail(email)
	if err != nil {
		utilities.RenderModalDialog(c, "Invalid email", "Please try again")
		return
	}

	templateFiles := []string{
		"root/layout.tmpl",
		"root/authentication/welcome.tmpl",
	}
	c.Status(http.StatusOK)
	utilities.RenderTemplate(c, "layout", gin.H{
		"Tenant": "MC",
		"Host":   "Madone Logistics",
		"Name": person.Name,
	}, templateFiles...)
}
func (app *PeopleController) ShowLogon(c *gin.Context) {
	// We need the layout and the specific welcome page.
	// The paths are relative to the 'templates' directory.
	templateFiles := []string{
		"root/layout.tmpl",
		"root/authentication/logon.tmpl",
	}

	// Call our custom renderer.
	// The name "layout.tmpl" tells the template engine which template definition to execute first.
	utilities.RenderTemplate(c, "layout", gin.H{
		"Tenant": "MC",
		"Host":   "Madone Logistics",
	}, templateFiles...)
}

func (app *PeopleController) HandleRegister(c *gin.Context) {
	var person people.Person

	// TODO add validate.validator
	name := c.PostForm("fullname")
	address := c.PostForm("address")
    email := c.PostForm("email")
	cell := c.PostForm("cell")
    password := c.PostForm("password")

	if email == "" {
		c.Status(http.StatusUnauthorized)
		utilities.RenderModalDialog(c, "No email provided", "Please try again")
		return
	}

	person, err := app.service.GetByEmail(email)
	if err == nil {
		c.Status(http.StatusExpectationFailed)
		utilities.RenderModalDialog(c, "Existing Email", "Please try again")
		return
	}
 
	if password == "" {
		c.Status(http.StatusUnauthorized)
		utilities.RenderModalDialog(c, "No password provided", "Please try again")
		return
	}

    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
		c.Status(http.StatusInternalServerError)
		utilities.RenderModalDialog(c, "Internal Server Error", "Please try again")
        return
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
		utilities.RenderModalDialog(c, "Internal Server Error,Could not create Person ", "Please try again")
        return
    }

	// This forces HTMX to reload the whole page without treating it as a fragment
	c.Status(http.StatusOK)
	c.Header("HX-Redirect", "/welcome/?email="+ person.Email)
}

func (app *PeopleController) HandleLogout(c *gin.Context) {

}

func (app *PeopleController) HandleNewPwd(c *gin.Context) {

}
// CurrentPerson retrieves the logged-in person from session (or nil).
func (app *PeopleController) CurrentPerson(c *gin.Context) *people.Person {
	if val, exists := c.Get(currentUserKey); exists {
		if u, ok := val.(*people.Person); ok {
			return u
		}
	}

	sess := sessions.Default(c)
	idVal := sess.Get(constants.PERSON_ID)
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

	var user people.Person
	user, err := app.service.GetByID(idUint)
	if err != nil {
		return nil
	}

	c.Set(currentUserKey, &user)
	return &user
}

func CheckPasswordHash(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

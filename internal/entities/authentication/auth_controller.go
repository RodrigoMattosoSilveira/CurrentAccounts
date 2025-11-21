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

func NewController(db *gorm.DB) *App {
	return &App{DB: db}
}

func (app *App) ShowLogin(c *gin.Context) {
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
func (app *App) HandleLogin(c *gin.Context) {

	var loginForm LoginForm
	if err := c.ShouldBind(&loginForm); err != nil {
		utilities.RenderModalDialog(c, "Invalid login form", "Please try again")
	}

	var person people.Person
	if err := app.DB.Where("email = ?", loginForm.Email).First(&person).Error; err != nil {
		utilities.RenderModalDialog(c, "Invalid email", "Please try again")
		return
	}

	if !CheckPasswordHash(person.Password, loginForm.Password) {
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
	c.Header("HX-Redirect", "/welcome/?email="+ person.Email)
	c.Status(http.StatusFound)
}
func (app *App) HandleWelcome(c *gin.Context) {
	email := c.Query("email")
	var person people.Person
	if err := app.DB.Where("email = ?",email).First(&person).Error; err != nil {
		utilities.RenderModalDialog(c, "Invalid email", "Please try again")
		return
	}

	templateFiles := []string{
		"root/layout.tmpl",
		"root/authentication/welcome.tmpl",
	}
	utilities.RenderTemplate(c, "layout", gin.H{
		"Tenant": "MC",
		"Host":   "Madone Logistics",
		"Name": person.Name,
	}, templateFiles...)
}
func (app *App) ShowLogon(c *gin.Context) {
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

func (app *App) HandleLogon(c *gin.Context) {

}

func (app *App) HandleLogoou(c *gin.Context) {

}

func (app *App) HandleLogout(c *gin.Context) {

}

func (app *App) HandleNewPwd(c *gin.Context) {

}
// CurrentPerson retrieves the logged-in person from session (or nil).
func (app *App) CurrentPerson(c *gin.Context) *people.Person {
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
	if err := app.DB.First(&user, idUint).Error; err != nil {
		return nil
	}
	c.Set(currentUserKey, &user)
	return &user
}

func CheckPasswordHash(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

package authentication

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/entities/people"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/utilities"
)


type LoginForm struct {
	Email string
	Password string
}

type Controller struct {
	service people.Service
}

func NewController(service people.Service) *Controller {
	return &Controller{service}
}

func (ctl *Controller) ShowLogin(c *gin.Context) {
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
func (ctl *Controller) HandleLogin(c *gin.Context) {

	var loginForm LoginForm
	if err := c.ShouldBind(&loginForm); err != nil {
		// HTMX partial snippet with hyperscript animation
		templateFileFN := utilities.GetTemplateFileFN("authentication/login_error.tmpl")
		c.HTML(http.StatusBadRequest, templateFileFN, gin.H{
			"msg": "Invalid login data",
		})
		return
	}

	var person people.Person
	person, err := ctl.service.GetByEmail(loginForm.Email) 
	if (err != nil) {
		// return c.Status(401).SendString("Invalid email")
		// c.HTML(http.StatusBadRequest, "person_invalid_email.tmpl", gin.H{"Error": "No user with this email"})
		templateFileFN := utilities.GetTemplateFileFN("authentication/login_error.tmpl")
		log.Printf("templateFileFN: %s", templateFileFN)		
		c.HTML(http.StatusBadRequest, templateFileFN, gin.H{
			"msg": "No user with this email",
		})
	}

	if !CheckPasswordHash(person.Password, loginForm.Password) {
		templateFileFN := utilities.GetTemplateFileFN("authentication/login_error.tmpl")
		log.Printf("templateFileFN: %s", templateFileFN)		
		c.HTML(http.StatusBadRequest, templateFileFN, gin.H{
			"msg": "Invalid password",
		})
		return
	}	

		// TODO implement this shortly
	// sess, err := ac.store.Get(c)
	// if err != nil {
	// 	// return c.Status(500).SendString("Session error")
	// 	messageLogin.Message = "Session error"
	// 	return c.Render("partials/auth/authMessage", messageLogin)
	// }

	// sess.Set("PersonId", person.ID)
	// sess.Set("PersonName", person.Name)

	// if err := sess.Save(); err != nil {
	// 	log.Printf("ERROR: Failed to save session: %v", err)
	// 	// return c.Status(500).SendString("Failed to save session")
	// 	messageLogin.Message = "Failed to save session"
	// 	return c.Render("partials/auth/authMessage", messageLogin)
	// }
	// c.Set("HX-Redirect", "/profile")

    // Generate JWT token
	// TODO implement this shortly
    // claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"username": person.Name,
	// 	"user_id": person.Email,
	// 	"role": "Associate",
   	// 	"iss": "ContasCorrentes",     
    //     "exp": 60 * 60 * 6 + time.Now().Unix(), // 6 hours from now
    // 	"iat": time.Now().Unix(),
    // })

    // secretKey := os.Getenv("JWT_KEY")
    // token, err := claims.SignedString([]byte(secretKey))
    // if err != nil {
	// 	// return c.SendStatus(fiber.StatusInternalServerError)
	// 	messageLogin.Message = "Fiber Internal Server Error"
	// 	return c.Render("partials/auth/authMessage", messageLogin)
    // }

	// // Create jwt cookie
	// cookie := new(fiber.Cookie)
	// cookie.Name = constants.COOKIE_NAME
	// cookie.Value =  token
	// cookie.MaxAge = 1000*60*60*6 // 6 hours
	// cookie.HTTPOnly = true
	// cookie.Secure = false
	// cookie.SameSite = "Secure"
	// c.Cookie(cookie)

	// This handler uses a different layout and its own template.
	templateFiles := []string{
		"root/simple_layout.tmpl",
		"root/hello/hello.tmpl",
	}

	// Execute the "simple_layout.tmpl" as the base.
	utilities.RenderTemplate(c, "simple_layout", gin.H{
		"Title": "Hello, Gin!",
    	"Body":  "Welcome to the Gin web framework.",
	}, templateFiles...)
}
func  (ctl *Controller) ShowLogon(c *gin.Context) {
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

func (ctl *Controller) HandleLogon(c *gin.Context) {

}

func (ctl *Controller) HandleLogoou(c *gin.Context) {

}

func (ctl *Controller) HandleLogout(c *gin.Context) {

}

func (ctl *Controller) HandleNewPwd(c *gin.Context) {

}
func CheckPasswordHash(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func renderReponseTemplate(c *gin.Context, layoutTpl string, layoutTplName string, partialTpl string, data gin.H) {
	// We need the layout and the specific welcome page.
	// The paths are relative to the 'templates' directory.
	templateFiles := []string{
		layoutTpl,
		partialTpl,
	}

	// Call our custom renderer.
	// The name "layout.tmpl" tells the template engine which template definition to execute first.
	utilities.RenderTemplate(c, layoutTplName, gin.H{
		"Tenant": "MC",
		"Host":   "Madone Logistics",
	}, templateFiles...)
}
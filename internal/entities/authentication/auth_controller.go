package authentication

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/constants"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/entities/people"
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

func checkPasswordHash(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

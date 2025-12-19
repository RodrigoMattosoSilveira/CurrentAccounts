package people

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/utilities"

)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service}
}

type personFormStruct struct {
	Fullname string
	Address string
	Cell string
	Email string
	Password string
}

type Atribute struct {
	DIV_ID string
	HX_URL string
	VALUE  string
}

var PersonEditForm []Atribute
// Displays all people
func (ctr *Controller) Index(c *fiber.Ctx) error {
	people, _ := ctr.service.GetAll()
	if len(people) != 0 {
		people = people[1:]
	}

	templates := utilities.NewTemplateStructures("top", "people/people.tmpl")
  	templates.AddData("person_row", "people/page/person_row.tmpl")

	templateData := utilities.NewTemplateData();
	templateData.AddData("People", people)
	c.Status(http.StatusOK)
	return utilities.RenderPage(c, templates.GetAllTemplateStructures(), templateData.GetAllData())
}

// Displays the form to register a new person
// A CRUD system may have two create routes for new users, one in authentication and 
// another in users, to allow for different scenarios and functionalities. The 
// authentication route would typically handle the process of creating a new user 
// account, ensuring that only authorized users can create new accounts. In contrast, the 
// users route would be used for creating user profiles or records that do not 
// necessarily require authentication. This separation helps in maintaining clarity and 
// organization in the API structure, making it easier for developers to understand and 
// manage the different aspects of user management within the application.
// 
// Therefor, for now, this route will redirect to the authentication logon page, until the
// requiements for having a separate new user creation page is clear.
func (ctr *Controller) New(c *fiber.Ctx) error  {
	route := "/logon"
	return c.Redirect(route, fiber.StatusSeeOther)
}

// Handles the submission of the new person form
// See  (ctr *Controller) New(c *fiber.Ctx) error for explanation
func (ctr *Controller) Create(c *fiber.Ctx) error  {
	route := "/register"
	return c.Redirect(route, fiber.StatusSeeOther)
}

// Displays a specific person
func (ctr *Controller) Show(c *fiber.Ctx) error{

	// get the id of the person to edit
	id, err := strconv.Atoi(c.Params("id"))
	// TODO find a better way to handle errors
	if err != nil {
		// invalid id
		files := []string {"root/layout.tmpl",  "person_index.tmpl",}
		data := map[string]any {"Tenant": "MC","Host":   "Madone Logistics", "Error": "invalid person id",}
		return utilities.RenderTemplateFiber(c, "layout", data, files ...)	
	}
	// get the person record
	person, err := ctr.service.GetByID(uint(id))
	if err != nil {
		// TODO find a better way to handle errors
		files := []string {"root/layout.tmpl",  "person_index.tmpl",}
		data := map[string]any {"Tenant": "MC","Host":   "Madone Logistics", "Error": "person record not found",}
		return utilities.RenderTemplateFiber(c, "layout", data, files ...)	
	}
	attributes := []Atribute{}
	attributes = append(attributes, buildAttribute("Name", person.Name, uint16(person.ID)))
	attributes = append(attributes, buildAttribute("Cell", person.Cell, uint16(person.ID)))
	attributes = append(attributes, buildAttribute("Email", person.Email, uint16(person.ID)))
	attributes = append(attributes, buildAttribute("Role", person.Role, uint16(person.ID)))

	// Default templates and data
	templates := utilities.NewTemplateStructures("top", "people/page/person_show.tmpl")
	templateData := utilities.NewTemplateData();
	templateData.AddData("ATTRIBUTES", attributes)
	c.Status(http.StatusOK)
	return utilities.RenderPage(c, templates.GetAllTemplateStructures(), templateData.GetAllData())
}

func (ctr *Controller) Update(c *fiber.Ctx)  error {
	// get the id of the person to edit
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		// invalid id
		files := []string {"root/layout.tmpl",  "person_index.tmpl",}
		data := map[string]any {"Tenant": "MC","Host":   "Madone Logistics", "Error": "invalid person id",}
		return utilities.RenderTemplateFiber(c, "layout", data, files ...)	
	}

	// get the person record
	person, err := ctr.service.GetByID(uint(id))
	if err != nil {
		files := []string {"root/layout.tmpl",  "person_index.tmpl",}
		data := map[string]any {"Tenant": "MC","Host":   "Madone Logistics", "Error": "person record not found",}
		return utilities.RenderTemplateFiber(c, "layout", data, files ...)	
	}

	// TODO Refactor to avoid code duplication with Create
	var form personFormStruct
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

    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
		c.Status(http.StatusInternalServerError)
		templateFiles := []string{"root/log.tmpl", "root/authentication/logon.tmpl", }
		return utilities.RenderTemplateFiber(c, "layout", map[string]any{
			"Tenant": "MC", "Host":   "Madone Logistics", "Error": "Invalid password, try again",
		}, templateFiles...)	
    }

    person = Person{
		Name: name,
		Address: address,
        Email:email,
		Cell: cell,
        Password: string(hash),
        Role: "Person",
    }
	err = ctr.service.Create(&person)
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

func (ctr *Controller) Delete(c *fiber.Ctx)  error {
	// get the id of the person to edit
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		// invalid id
		files := []string {"root/layout.tmpl",  "person_index.tmpl",}
		data := map[string]any {"Tenant": "MC","Host":   "Madone Logistics", "Error": "invalid person id",}
		return utilities.RenderTemplateFiber(c, "layout", data, files ...)	
	}

	// get the person record
	_, err = ctr.service.GetByID(uint(id))
	if err != nil {
		files := []string {"root/layout.tmpl",  "person_index.tmpl",}
		data := map[string]any {"Tenant": "MC","Host":   "Madone Logistics", "Error": "person record not found",}
		return utilities.RenderTemplateFiber(c, "layout", data, files ...)	
	}

	// delete the person record
	err = ctr.service.Delete(uint(id))
	if err != nil {
		files := []string {"root/layout.tmpl",  "person_index.tmpl",}
		data := map[string]any {"Tenant": "MC","Host":   "Madone Logistics", "Error": "unable to delete person record",}
		return utilities.RenderTemplateFiber(c, "layout", data, files ...)	
	}

	return c.Redirect("/people]", fiber.StatusOK)
}
func buildAttribute(attributeName string, attributeValue string, personID uint16) Atribute {
	return Atribute {DIV_ID: attributeName, HX_URL: "/person/" + strconv.Itoa(int(personID)) + "/edit", VALUE:  attributeValue,}
}

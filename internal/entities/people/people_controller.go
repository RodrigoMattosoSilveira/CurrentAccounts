package people

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"

	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/utilities"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service}
}

type Atribute struct {
	DIV_ID string
	HX_URL string
	VALUE  string
	ERROR  string
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

	templateData := utilities.NewTemplateData()
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
func (ctr *Controller) New(c *fiber.Ctx) error {
	route := "/logon"
	return c.Redirect(route, fiber.StatusSeeOther)
}

// Handles the submission of the new person form
// See  (ctr *Controller) New(c *fiber.Ctx) error for explanation
func (ctr *Controller) Create(c *fiber.Ctx) error {
	route := "/register"
	return c.Redirect(route, fiber.StatusSeeOther)
}

// Displays a specific person
func (ctr *Controller) Show(c *fiber.Ctx) error {

	// get the id of the person to edit
	id, err := strconv.Atoi(c.Params("id"))
	// TODO find a better way to handle errors
	if err != nil {
		// invalid id
		files := []string{"root/layout.tmpl", "person_index.tmpl"}
		data := map[string]any{"Tenant": "MC", "Host": "Madone Logistics", "Error": "invalid person id"}
		return utilities.RenderTemplateFiber(c, "layout", data, files...)
	}
	// get the person record
	person, err := ctr.service.GetByID(uint(id))
	if err != nil {
		// TODO find a better way to handle errors
		files := []string{"root/layout.tmpl", "person_index.tmpl"}
		data := map[string]any{"Tenant": "MC", "Host": "Madone Logistics", "Error": "person record not found"}
		return utilities.RenderTemplateFiber(c, "layout", data, files...)
	}
	/*
			<div id="{{.DIV_ID}}" class="form-row-grid mb-2">
				<label class="form-label fw-semibold mb-0">
				{{.DIV_ID}}:
			</label>
			<span
				class="editable text-primary"
				hx-get="{{.HX_URL}}"
				hx-target="{{.DIV_ID}}"
				hx-swap="outerHTML"
				style="cursor:pointer"
			>
				{{.VALUE}}
			</span>
		</div>
	*/
	attributes := []Atribute{}
	attributes = append(attributes, buildAttribute("Name", person.Name, uint16(person.ID)))
	attributes = append(attributes, buildAttribute("Cell", person.Cell, uint16(person.ID)))
	attributes = append(attributes, buildAttribute("Email", person.Email, uint16(person.ID)))
	attributes = append(attributes, buildAttribute("Role", person.Role, uint16(person.ID)))

	// Default templates and data
	templates := utilities.NewTemplateStructures("top", "people/page/person_show.tmpl")
	templates.AddData("partial", "people/partial/person_edit_partial.tmpl")
	templateData := utilities.NewTemplateData()
	templateData.AddData("ATTRIBUTES", attributes)
	c.Status(http.StatusOK)
	return utilities.RenderPage(c, templates.GetAllTemplateStructures(), templateData.GetAllData())
}
func (ctr *Controller) EditName(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	person, _ := ctr.service.GetByID(uint(id))

	data := Atribute{
		DIV_ID: "Name",
		HX_URL: "/person/" + strconv.Itoa(int(person.ID)) + "/name",
		VALUE:  person.Name,
		ERROR:  "",
	}
	return c.Render("person_update_partial", data)
}
func (ctr *Controller) UpdateName(c *fiber.Ctx) error {
	type formStruct struct {
		Name string
	}

	id, _ := strconv.Atoi(c.Params("id"))
	person, _ := ctr.service.GetByID(uint(id))

	data := fiber.Map{
		"DIV_ID": "Name",
		"HX_URL": "/person/" + strconv.Itoa(int(person.ID)) + "/name",
		"VALUE":  person.Name,
		"ERROR":  "",
	}

	var form formStruct
	if err := c.BodyParser(&form); err != nil {
		data["ERROR"] = "Unable to parse name from form, try again"
		return c.Render("person_update_partial", data)
	}

	name := form.Name
	validate := validator.New()
	err := validate.Var(name, "required,min=2,max=40")
	if err != nil {
		data["ERROR"] = "Invalid name, required,min=2,max=40. Try again"
		return c.Render("person_update_partial", data)
	}
	person.Name = name
	data["VALUE"] = name
	err = ctr.service.Update(&person)
	if err != nil {
		data["ERROR"] = "Unable to update name, try again"
		return c.Render("person_update_partial", data)
	}

	return utilities.RenderPartial(c, "people/partial/person_edit_partial.tmpl", data)
}
func (ctr *Controller) EditAddress(c *fiber.Ctx) error {
	return nil
}
func (ctr *Controller) UpdateAddress(c *fiber.Ctx) error {
	return nil
}
func (ctr *Controller) EditEmail(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	person, _ := ctr.service.GetByID(uint(id))

	data := Atribute{
		DIV_ID: "Email",
		HX_URL: "/person/" + strconv.Itoa(int(person.ID)) + "/email",
		VALUE:  person.Email,
		ERROR:  "",
	}
	return c.Render("person_update_partial", data)
}
func (ctr *Controller) UpdateEmail(c *fiber.Ctx) error {
	type formStruct struct {
		Email string
	}

	id, _ := strconv.Atoi(c.Params("id"))
	person, _ := ctr.service.GetByID(uint(id))

	data := fiber.Map{
		"DIV_ID": "Email",
		"HX_URL": "/person/" + strconv.Itoa(int(person.ID)) + "/email",
		"VALUE":  person.Email,
		"ERROR":  "",
	}

	var form formStruct
	if err := c.BodyParser(&form); err != nil {
		data["ERROR"] = "Unable to parse email from form, try again"
		return c.Render("person_update_partial", data)
	}

	email := form.Email
	data["VALUE"] = email
	validate := validator.New()
	err := validate.Var(email, "required,email")
	if err != nil {
		data["ERROR"] = "Invalid email, required,email. Try again"
		return c.Render("person_update_partial", data)
	}
	person.Email = email
	err = ctr.service.Update(&person)
	if err != nil {
		data["ERROR"] = "Unable to update email, try again"
		return c.Render("person_update_partial", data)
	}

	return utilities.RenderPartial(c, "people/partial/person_edit_partial.tmpl", data)
}
func (ctr *Controller) EditCell(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	person, _ := ctr.service.GetByID(uint(id))

	data := fiber.Map{
		"DIV_ID": "Cell",
		"HX_URL": "/person/" + strconv.Itoa(int(person.ID)) + "/cell",
		"VALUE":  person.Cell,
		"ERROR":  "",
	}
	return c.Render("person_update_partial", data)
}
func (ctr *Controller) UpdateCell(c *fiber.Ctx) error {
	type formStruct struct {
		Cell string
	}

	id, _ := strconv.Atoi(c.Params("id"))
	person, _ := ctr.service.GetByID(uint(id))

	data := fiber.Map{
		"DIV_ID": "Cell",
		"HX_URL": "/person/" + strconv.Itoa(int(person.ID)) + "/cell",
		"VALUE":  person.Cell,
		"ERROR":  "",
	}

	var form formStruct
	if err := c.BodyParser(&form); err != nil {
		data["ERROR"] = "Unable to parse cell from form, try again"
		return c.Render("person_update_partial", data)
	}

	cell := form.Cell
	validate := validator.New()
	err := validate.Var(cell, "required,min=9,max=12")
	if err != nil {
		data["ERROR"] = "Invalid cell, required,min=9,max=12. Try again"
		return c.Render("person_update_partial", data)
	}
	person.Cell = cell
	data["VALUE"] = cell
	err = ctr.service.Update(&person)
	if err != nil {
		data["ERROR"] = "Unable to update cell, try again"
		return c.Render("person_update_partial", data)
	}

	return utilities.RenderPartial(c, "people/partial/person_edit_partial.tmpl", data)
}
func (ctr *Controller) EditRole(c *fiber.Ctx) error {			
	id, _ := strconv.Atoi(c.Params("id"))
	person, _ := ctr.service.GetByID(uint(id))

	data := fiber.Map{
		"DIV_ID": "Role",
		"HX_URL": "/person/" + strconv.Itoa(int(person.ID)) + "/role",
		"VALUE":  person.Role,
		"ERROR":  "",
	}
	return c.Render("person_update_partial", data)
}		
func (ctr *Controller) UpdateRole(c *fiber.Ctx) error {
	type formStruct struct {
		Role string
	}

	id, _ := strconv.Atoi(c.Params("id"))
	person, _ := ctr.service.GetByID(uint(id))

	data := fiber.Map{
		"DIV_ID": "Role",
		"HX_URL": "/person/" + strconv.Itoa(int(person.ID)) + "/role",
		"VALUE":  person.Role,
		"ERROR":  "",
	}

	var form formStruct
	if err := c.BodyParser(&form); err != nil {
		data["ERROR"] = "Unable to parse role from form, try again"
		return c.Render("person_update_partial", data)
	}

	// TODO add validate.validator
	role := form.Role

	validate := validator.New()
	err := validate.Var(role, "required,oneof=Person Operator Application Tenant System")
	if err != nil {
		data["ERROR"] = "Invalid role, required,oneof=Person Operator Application Tenant System Try again"
		return c.Render("person_update_partial", data)
	}
	person.Role = role
	data["VALUE"] = role
	err = ctr.service.Update(&person)
	if err != nil {
		data["ERROR"] = "Unable to update role, try again"
		return c.Render("person_update_partial", data)
	}

	return utilities.RenderPartial(c, "people/partial/person_edit_partial.tmpl", data)
}

func (ctr *Controller) Delete(c *fiber.Ctx) error {
	// get the id of the person to edit
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		// invalid id
		files := []string{"root/layout.tmpl", "person_index.tmpl"}
		data := map[string]any{"Tenant": "MC", "Host": "Madone Logistics", "Error": "invalid person id"}
		return utilities.RenderTemplateFiber(c, "layout", data, files...)
	}

	// get the person record
	_, err = ctr.service.GetByID(uint(id))
	if err != nil {
		files := []string{"root/layout.tmpl", "person_index.tmpl"}
		data := map[string]any{"Tenant": "MC", "Host": "Madone Logistics", "Error": "person record not found"}
		return utilities.RenderTemplateFiber(c, "layout", data, files...)
	}

	// delete the person record
	err = ctr.service.Delete(uint(id))
	if err != nil {
		files := []string{"root/layout.tmpl", "person_index.tmpl"}
		data := map[string]any{"Tenant": "MC", "Host": "Madone Logistics", "Error": "unable to delete person record"}
		return utilities.RenderTemplateFiber(c, "layout", data, files...)
	}

	return c.Redirect("/people]", fiber.StatusOK)
}
func buildAttribute(attributeName string, attributeValue string, personID uint16) Atribute {
	return Atribute{
		DIV_ID: attributeName,
		HX_URL: "/person/" + strconv.Itoa(int(personID)) + "/" + attributeName,
		VALUE:  attributeValue,
	}
}

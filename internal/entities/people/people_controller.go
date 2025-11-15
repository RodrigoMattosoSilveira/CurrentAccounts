package people

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service}
}

func (ctr *Controller) Index(c *gin.Context) {
	people, _ := ctr.service.GetAll()
	c.HTML(http.StatusOK, "person_index.tmpl", gin.H{"People": people})
}

func (ctr *Controller) New(c *gin.Context) {
	c.HTML(http.StatusOK, "person_new.tmpl", nil)
}

func (ctr *Controller) Create(c *gin.Context) {
	var form Person
	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusBadRequest, "person_new.tmpl", gin.H{"Error": "Invalid form data"})
		return
	}
	ctr.service.Create(&form)
	c.Redirect(http.StatusFound, "/people")
}

func (ctr *Controller) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	person, _ := ctr.service.GetByID(uint(id))
	c.HTML(http.StatusOK, "person_show.tmpl", gin.H{"Person": person})
}

func (ctr *Controller) Edit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	person, _ := ctr.service.GetByID(uint(id))
	c.HTML(http.StatusOK, "person_edit.tmpl", gin.H{"Person": person})
}

func (ctr *Controller) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	person, _ := ctr.service.GetByID(uint(id))

	if err := c.ShouldBind(&person); err != nil {
		c.HTML(http.StatusBadRequest, "person_edit.tmpl", gin.H{"Error": "Invalid data"})
		return
	}

	ctr.service.Update(&person)
	c.Redirect(http.StatusFound, "/people")
}

func (ctr *Controller) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ctr.service.Delete(uint(id))
	c.Redirect(http.StatusFound, "/people")
}

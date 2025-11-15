package people

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Person, error)
	FindByID(id uint) (Person, error)
	FindByEmail(email string) (Person, error)
	Create(p *Person) error
	Update(p *Person) error
	Delete(p *Person) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Person, error) {
	var list []Person
	err := r.db.Find(&list).Error
	return list, err
}

func (r *repository) FindByID(id uint) (Person, error) {
	var p Person
	err := r.db.First(&p, id).Error
	return p, err
}

func (r *repository) FindByEmail(emailToFind string) (Person, error) {
	var person Person

	error := r.db.Where("email = ?", emailToFind).First(&person).Error
	return person, error
}

func (r *repository) Create(p *Person) error {
	return r.db.Create(p).Error
}

func (r *repository) Update(p *Person) error {
	return r.db.Save(p).Error
}

func (r *repository) Delete(p *Person) error {
	return r.db.Delete(p).Error
}

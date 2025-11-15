package people

import "gorm.io/gorm"

type Person struct {
    gorm.Model
	Name     string `form:"name" validate:"required,min=2"`
	Email    string `form:"email" gorm:"unique" validate:"required,email"`
	Cell     string `form:"cell" validate:"required,min=9"`
	Password string `form:"password" validate:"required,min=8,max=12"`
	Role     string `form:"role"`
	// Rg string `gorm:"unique"`
	// Cpf string `gorm:"unique"`
	// Street string
	// District string
	// City string
	// Cep string `form:"CEP" validate:"required,cepx"`
	// State string

	// Bank string
	// BankNumber string
	// Account string
	// Pix string `form:"PIX" gorm:"unique" validate:"required,pix"`
	// EmergencyName string `form:"name" validate:"required,min=2"`
	// EmergencyEmail string `form:"email" validate:"required,email"`
	// EmergencyCell string
}
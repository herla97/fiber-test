package models

import (
	"errors"
	"fiapi/utils"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"gorm.io/gorm"
)

// User struct
type User struct {
	Base
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"password,omitempty"`
	// Lastname1 string `gorm:"not null" json:"lastname1"`
	// Lastname2 string `json:"lastname2"`
}

// Exist user function. TODO: Busqueda dinámica.
func (u *User) Exist(tx *gorm.DB) bool {
	err := tx.First(u, "email = ?", u.Email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

// Validate user struture
// See more: https://github.com/go-ozzo/ozzo-validation
func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name,
			validation.Required,
			validation.Length(5, 20).Error("mínimo de 5 caracteres"),
		),
		validation.Field(&u.Email,
			validation.Required.Error("no puede ser vacío"),
			is.Email.Error("formato inválido"),
		),
		validation.Field(&u.Password,
			validation.Required.Error("no puede ser vacío"),
			validation.Length(5, 50),
		),
	)
}

// BeforeCreate Hook GORM function
// See more: https://gorm.io/docs/hooks.html#Creating-an-object
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Validate user struture
	if err := u.Validate(); err != nil {
		return err
	}

	// Validate user exist
	if u.Exist(tx) {
		return errors.New("usuario existente")
	}

	u.Password, err = utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}

// AfterCreate Hook GORM function
// See more: https://gorm.io/docs/hooks.html#Creating-an-object
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	// Clean password
	u.Password = ""
	return nil
}

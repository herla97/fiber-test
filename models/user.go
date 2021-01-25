package models

import "golang.org/x/crypto/bcrypt"

// User struct
type User struct {
	Base
	Username  string `gorm:"unique_index;not null" json:"username"`
	Email     string `gorm:"unique_index;not null" json:"email"`
	Password  string `gorm:"not null" json:"password,omitempty"`
	Name      string `gorm:"not null" json:"name"`
	Lastname1 string `gorm:"not null" json:"lastname1"`
	Lastname2 string `json:"lastname2"`
}

// HashPassword User Hash Password function
func (u *User) HashPassword() error {
	ph, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(ph)
	return nil
}

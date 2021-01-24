package models

// User struct
type User struct {
	Base
	Username  string `gorm:"unique_index;not null" json:"username"`
	Email     string `gorm:"unique_index;not null" json:"email"`
	Password  string `gorm:"not null" json:"password"`
	Name      string `gorm:"not null" json:"name"`
	Lastname1 string `gorm:"not null" json:"lastname1"`
	Lastname2 string `json:"lastname2"`
}

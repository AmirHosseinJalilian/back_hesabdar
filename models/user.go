package models

type User struct {
	ID       int    `gorm:"primary_key"`
	Username string `gorm:"unique"`
	Password string
	// Add other fields as needed
}

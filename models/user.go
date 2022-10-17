package models

type User struct {
	ID           int    `json:"id" gorm:"primary_key"`
	Email        string `json:"email" gorm:"unique"`
	Password     string `json:"password"`
	Full_name    string `json:"full_name"`
	Phone_number int    `json:"phone_number"`
	Gender       string `json:"gender"`
}

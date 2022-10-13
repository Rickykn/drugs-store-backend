package models

type User struct {
	ID           int    `json:"id" gorm:"primary_key"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Full_name    string `json:"full_name"`
	Phone_number int    `json:"phone_number"`
	Address      string `json:"address"`
	City         string `json:"city"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	Gender       string `json:"gender"`
}

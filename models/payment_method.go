package models

type PaymentMethod struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

package models

type Courier struct {
	ID                  int    `json:"id" gorm:"primary_key"`
	Price               int    `json:"price"`
	Estimate_deliveried string `json:"estimate_deliveried"`
}

package models

import "time"

type Transaction struct {
	ID                 int       `json:"id" gorm:"primary_key"`
	Total_price        int       `json:"total_price"`
	Status_transaction int       `json:"status_transaction"`
	Transaction_date   time.Time `json:"transaction_date"`
	Prescription_image string    `json:"prescription_image"`
	Address_id         *int      `json:"address_id"`
	Address            Address   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Courier_id         *int      `json:"courier_id"`
	Courier            Courier   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

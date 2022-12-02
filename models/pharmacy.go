package models

import "time"

type Pharmacy struct {
	ID               int       `json:"id" gorm:"primary_key"`
	Pharmacy_name    string    `json:"pharmacy_name" gorm:"unique"`
	Pharmacy_address string    `json:"pharmacy_address"`
	City             string    `json:"city"`
	Latitude         string    `json:"latitude"`
	Longitude        string    `json:"longitude"`
	License_number   int32     `json:"license_number"`
	Phone_number     int32     `json:"phone_number"`
	Operation_start  time.Time `json:"operation_start"`
	Operation_end    time.Time `json:"operation_end"`
}

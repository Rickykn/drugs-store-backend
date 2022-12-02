package models

import "time"

type Voucher struct {
	ID              int       `json:"id" gorm:"primary_key"`
	Voucher_name    string    `json:"voucher_name"`
	Quantity        int       `json:"quantty"`
	Expired_date    time.Time `json:"expired_date"`
	Created_date    time.Time `json:"Created_date"`
	Discount_amount int       `json:"discount_amount"`
}

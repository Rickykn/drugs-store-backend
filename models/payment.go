package models

type Payment struct {
	ID                int           `json:"id" gorm:"primary_key"`
	Amount            int           `json:"amount"`
	Payment_method_id *int          `json:"payment_method_id"`
	PaymentMethod     PaymentMethod `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Transaction_id    *int          `json:"transaction_id"`
	Transaction       Transaction   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Payment_image     string        `json:"payment_image"`
}

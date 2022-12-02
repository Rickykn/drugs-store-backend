package models

type TransactionItem struct {
	ID             int         `json:"id" gorm:"primary_key"`
	Inventory_id   *int        `json:"inventory_id"`
	Inventory      Inventory   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Transaction_id *int        `json:"transaction_id"`
	Transaction    Transaction `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Price          int         `json:"price"`
	Quantity       int         `json:"quantity"`
}

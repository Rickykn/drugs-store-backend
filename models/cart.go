package models

type Cart struct {
	ID           int       `json:"id" gorm:"primary_key"`
	Quantity     int       `json:"quantity"`
	Total_price  int       `json:"total_price"`
	User_id      *int      `json:"user_id"`
	User         User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Inventory_id *int      `json:"inventory_id"`
	Inventory    Inventory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

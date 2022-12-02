package models

type Address struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Address   string `json:"address"`
	City      string `json:"city"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	User_id   *int   `json:"user_id"`
	User      User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

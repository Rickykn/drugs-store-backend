package models

type Inventory struct {
	ID          int      `json:"id" gorm:"primary_key"`
	Stock       int      `json:"stock"`
	Pharmacy_id *int     `json:"pharmacy_id"`
	Pharmacy    Pharmacy `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Drug_id     *int     `json:"drug_id"`
	Drug        Drug     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

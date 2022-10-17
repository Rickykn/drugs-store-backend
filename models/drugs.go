package models

type Drug struct {
	ID                  int    `json:"id" gorm:"primary_key"`
	Name                string `json:"name"`
	Generic_name        string `json:"generic_name"`
	Content             string `json:"content"`
	Manufacture         string `json:"manufacture"`
	Description         string `json:"description"`
	Drug_classification string `json:"drug_classification"`
	Drug_form           string `json:"drug_form"`
	Unit_in_pack        string `json:"unit_in_pack"`
	Selling_unit        string `json:"selling_unit"`
	Weight              int    `json:"weight"`
	Height              int    `json:"height"`
	Length              int    `json:"length"`
	Widht               int    `json:"widht"`
	Image               string `json:"image"`
}

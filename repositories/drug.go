package repositories

import (
	"github.com/Rickykn/drugs-store-backend.git/dtos"
	"github.com/Rickykn/drugs-store-backend.git/models"
	"gorm.io/gorm"
)

type DrugRepository interface {
	Create(inputDrug *dtos.CreateDrugDTO) (*models.Drug, error)
}

type drugRepository struct {
	db *gorm.DB
}

type DRConfig struct {
	DB *gorm.DB
}

func NewDrugRepository(c *DRConfig) DrugRepository {
	return &drugRepository{
		db: c.DB,
	}
}

func (d *drugRepository) Create(inputDrug *dtos.CreateDrugDTO) (*models.Drug, error) {

	newDrug := &models.Drug{
		Name:                inputDrug.Name,
		Generic_name:        inputDrug.Generic_name,
		Content:             inputDrug.Content,
		Manufacture:         inputDrug.Manufacture,
		Description:         inputDrug.Description,
		Drug_classification: inputDrug.Drug_classification,
		Drug_form:           inputDrug.Drug_form,
		Unit_in_pack:        inputDrug.Unit_in_pack,
		Selling_unit:        inputDrug.Selling_unit,
		Weight:              inputDrug.Weight,
		Height:              inputDrug.Height,
		Widht:               inputDrug.Widht,
	}

	result := d.db.Create(&newDrug)
	return newDrug, result.Error
}

package repositories

import "gorm.io/gorm"

type PharmacyRepository interface {
	Create()
}

type pharmacyRepository struct {
	db *gorm.DB
}

type PRConfig struct {
	DB *gorm.DB
}

func NewPharmacyRepository(c *PRConfig) PharmacyRepository {
	return &pharmacyRepository{
		db: c.DB,
	}
}

func (p *pharmacyRepository) Create() {

}

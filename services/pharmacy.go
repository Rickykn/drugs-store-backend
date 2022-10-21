package services

import (
	r "github.com/Rickykn/drugs-store-backend.git/repositories"
)

type PharmacyService interface {
	CreateNewPharmacy()
}

type pharmacyService struct {
	pharmacyRepository r.PharmacyRepository
}

type PSConfig struct {
	PharmacyRepository r.PharmacyRepository
}

func NewPharmacyService(c *PSConfig) PharmacyService {
	return &pharmacyService{
		pharmacyRepository: c.PharmacyRepository,
	}
}

func (p *pharmacyService) CreateNewPharmacy() {

}

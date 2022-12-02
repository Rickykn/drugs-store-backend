package services

import (
	"github.com/Rickykn/drugs-store-backend.git/dtos"
	help "github.com/Rickykn/drugs-store-backend.git/helpers"
	"github.com/Rickykn/drugs-store-backend.git/models"
	r "github.com/Rickykn/drugs-store-backend.git/repositories"
)

type DrugService interface {
	CreateDrug(drugInput *dtos.CreateDrugDTO) (*models.Drug, *help.JsonResponse, error)
}

type drugService struct {
	drugRepository r.DrugRepository
}

type DSConfig struct {
	DrugRepository r.DrugRepository
}

func NewDrugService(c *DSConfig) DrugService {
	return &drugService{
		drugRepository: c.DrugRepository,
	}
}

func (d *drugService) CreateDrug(drugInput *dtos.CreateDrugDTO) (*models.Drug, *help.JsonResponse, error) {
	newDrug, err := d.drugRepository.Create(drugInput)

	if err != nil {
		return nil, help.HandlerError(500, "Server Error", nil), err
	}

	return newDrug, help.HandlerSuccess(201, "Success Create New Drug", newDrug), err
}

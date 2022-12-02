package handlers

import (
	"net/http"

	"github.com/Rickykn/drugs-store-backend.git/dtos"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateNewDrug(c *gin.Context) {
	var drugInput *dtos.CreateDrugDTO
	err := c.ShouldBindJSON(&drugInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	drug, response, _ := h.drugService.CreateDrug(drugInput)

	if response.Error {
		c.JSON(response.Code, gin.H{
			"message":     response.Message,
			"status code": response.Code,
			"data":        response.Data,
		})
	} else {
		newDrug := &dtos.CreateDrugDTO{
			Name:                drug.Name,
			Generic_name:        drug.Generic_name,
			Content:             drug.Content,
			Manufacture:         drug.Manufacture,
			Description:         drug.Description,
			Drug_classification: drug.Drug_classification,
			Drug_form:           drug.Drug_form,
			Unit_in_pack:        drug.Unit_in_pack,
			Selling_unit:        drug.Selling_unit,
			Weight:              drug.Weight,
			Height:              drug.Height,
			Length:              drug.Length,
		}

		c.JSON(response.Code, gin.H{
			"message":     response.Message,
			"status code": response.Code,
			"data":        newDrug,
		})
	}
}

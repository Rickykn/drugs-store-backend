package handlers

import (
	"net/http"

	"github.com/Rickykn/drugs-store-backend.git/dtos"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterUser(c *gin.Context) {
	var registerInput *dtos.UserRegisterDTO
	err := c.ShouldBindJSON(&registerInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return

	}

	user, response, _ := h.userService.Register(registerInput)

	if response.Error {
		c.JSON(response.Code, gin.H{
			"message":     response.Message,
			"status code": response.Code,
			"data":        response.Data,
		})
	} else {
		newUser := &dtos.UserRegisterResponse{
			Name:         user.Full_name,
			Email:        user.Email,
			Phone_number: user.Phone_number,
			Gender:       user.Gender,
		}

		c.JSON(response.Code, gin.H{
			"message":     response.Message,
			"status code": response.Code,
			"data":        newUser,
		})
	}
}

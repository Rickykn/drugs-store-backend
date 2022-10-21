package main

import (
	"fmt"
	"log"

	"github.com/Rickykn/drugs-store-backend.git/database"
	"github.com/Rickykn/drugs-store-backend.git/handlers"
	"github.com/Rickykn/drugs-store-backend.git/repositories"
	"github.com/Rickykn/drugs-store-backend.git/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("E-Wallet backend")

	errEnv := godotenv.Load()

	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	errConnect := database.Connect()

	if errConnect != nil {
		panic(errConnect)
	}

	r := gin.Default()
	ur := repositories.NewUserRepository(&repositories.URConfig{
		DB: database.Get(),
	})
	us := services.NewUserService(&services.USConfig{
		UserRepository: ur,
	})

	dr := repositories.NewDrugRepository(&repositories.DRConfig{
		DB: database.Get(),
	})

	ds := services.NewDrugService(&services.DSConfig{
		DrugRepository: dr,
	})

	h := handlers.New(&handlers.HandlerConfig{

		UserService: us,
		DrugService: ds,
	})

	users := r.Group("/users")
	{

		users.POST("/register", h.RegisterUser)
		users.POST("/login", h.LoginUser)
	}

	drugs := r.Group("/drugs")
	{
		drugs.POST("/", h.CreateNewDrug)
	}

	err := r.Run()

	if err != nil {
		panic(err)
	}
}

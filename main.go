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

	h := handlers.New(&handlers.HandlerConfig{

		UserService: us,
	})

	users := r.Group("/users")
	{

		users.POST("/register", h.RegisterUser)
	}

	err := r.Run()

	if err != nil {
		panic(err)
	}
}

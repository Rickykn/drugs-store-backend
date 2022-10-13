package main

import (
	"fmt"
	"log"

	"github.com/Rickykn/drugs-store-backend.git/database"
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
	err := r.Run()

	if err != nil {
		panic(err)
	}
}

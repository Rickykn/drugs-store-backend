package database

import (
	"fmt"
	"os"

	"github.com/Rickykn/drugs-store-backend.git/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() (err error) {

	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORTDB")
	DB_USER := os.Getenv("DBUSER")
	DB_PASS := os.Getenv("DBPASSWORD")
	DB_NAME := os.Getenv("DBNAME")

	fmt.Println(DB_NAME)

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		HOST, DB_USER, DB_PASS, DB_NAME, PORT)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Address{})
	db.AutoMigrate(&models.Drug{})
	db.AutoMigrate(&models.Inventory{})
	db.AutoMigrate(&models.Pharmacy{})
	db.AutoMigrate(&models.Cart{})
	db.AutoMigrate(&models.Courier{})
	db.AutoMigrate(&models.Transaction{})
	db.AutoMigrate(&models.TransactionItem{})
	db.AutoMigrate(&models.Payment{})
	db.AutoMigrate(&models.PaymentMethod{})
	db.AutoMigrate(&models.Voucher{})

	return nil
}

func Get() *gorm.DB {
	return db
}

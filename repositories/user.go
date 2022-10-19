package repositories

import (
	"github.com/Rickykn/drugs-store-backend.git/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(email, name, password, gender string, phone_number int) (*models.User, error)
	FindOneUser(email string) (*models.User, int, error)
}

type userRepository struct {
	db *gorm.DB
}

type URConfig struct {
	DB *gorm.DB
}

func NewUserRepository(c *URConfig) UserRepository {
	return &userRepository{
		db: c.DB,
	}
}

func (u *userRepository) CreateUser(email, name, password, gender string, phone_number int) (*models.User, error) {
	newUser := &models.User{
		Full_name:    name,
		Email:        email,
		Password:     password,
		Phone_number: phone_number,
		Gender:       gender,
	}

	result := u.db.Create(&newUser)

	return newUser, result.Error
}

func (u *userRepository) FindOneUser(email string) (*models.User, int, error) {
	var user *models.User

	result := u.db.Where("email = ?", email).First(&user)

	return user, int(result.RowsAffected), result.Error
}

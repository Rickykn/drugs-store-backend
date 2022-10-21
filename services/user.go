package services

import (
	"github.com/Rickykn/drugs-store-backend.git/dtos"
	help "github.com/Rickykn/drugs-store-backend.git/helpers"
	"github.com/Rickykn/drugs-store-backend.git/models"
	r "github.com/Rickykn/drugs-store-backend.git/repositories"
)

type UserService interface {
	Register(registerInput *dtos.UserRegisterDTO) (*models.User, *help.JsonResponse, error)
	Login(loginInput *dtos.LoginUserDTO) (*dtos.TokenDTO, *help.JsonResponse, error)
}
type userService struct {
	userRepository r.UserRepository
}

type USConfig struct {
	UserRepository r.UserRepository
}

func NewUserService(c *USConfig) UserService {

	return &userService{
		userRepository: c.UserRepository,
	}
}

func (u *userService) Register(registerInput *dtos.UserRegisterDTO) (*models.User, *help.JsonResponse, error) {
	_, row, err := u.userRepository.FindOneUser(registerInput.Email)

	if row == 1 {
		return nil, help.HandlerError(400, "Email has been Taken", nil), err
	}

	hasingPassword, _ := help.HashPassword(registerInput.Password)
	newUser, err := u.userRepository.CreateUser(registerInput.Email, registerInput.Name, hasingPassword, registerInput.Gender, registerInput.Phone_number)

	if err != nil {
		return nil, help.HandlerError(500, "Server Error", nil), err
	}

	return newUser, help.HandlerSuccess(201, "Success register account", newUser), nil
}
func (u *userService) Login(loginInput *dtos.LoginUserDTO) (*dtos.TokenDTO, *help.JsonResponse, error) {

	findUser, row, err := u.userRepository.FindOneUser(loginInput.Email)
	if row == 0 || err != nil {
		return nil, help.HandlerError(404, "Wrong email or password", nil), err
	}

	isPasswordCorrect := help.CheckPasswordHash(loginInput.Password, findUser.Password)
	if !isPasswordCorrect {
		return nil, help.HandlerError(404, "Wrong email or password", nil), err
	}
	dataToken := &dtos.ResponseTokenDTO{
		ID:    findUser.ID,
		Name:  findUser.Full_name,
		Email: findUser.Email,
	}
	tokenString, err := help.CreateJwt(dataToken)

	if err != nil {
		return nil, help.HandlerError(500, "Server Error", nil), err
	}

	return &dtos.TokenDTO{IDToken: tokenString}, help.HandlerSuccess(200, "Login success", &dtos.TokenDTO{IDToken: tokenString}), nil
}

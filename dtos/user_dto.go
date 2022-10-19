package dtos

type UserRegisterDTO struct {
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Phone_number int    `json:"phone_number" binding:"required"`
	Gender       string `json:"gender" binding:"required"`
}

type UserRegisterResponse struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone_number int    `json:"phone_number" `
	Gender       string `json:"gender" `
}

type LoginUserDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenDTO struct {
	IDToken string `json:"idToken"`
}

type ResponseTokenDTO struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

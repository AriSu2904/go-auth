package dto

type RegisterUserInput struct {
	Email    string `json:"email" binding:"required"`
	Persona  string `json:"persona" binding:"required"`
	Password string `json:"password" binding:"required"`
}

package dto

type RegisterUserInput struct {
	Email    string `json:"email" binding:"required"`
	Persona  string `json:"persona" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserInput struct {
	UniqueId string `json:"uniqueId" binding:"required"` // can be email or persona
	Password string `json:"password" binding:"required"`
}

type AdditionalHeader struct {
	DeviceId   string
	DeviceInfo string
}

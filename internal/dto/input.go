package dto

type RegisterUserInput struct {
	Email    string `validate:"email" json:"email"`
	Persona  string `validate:"alphanum,min=3,max=30" json:"persona"`
	Password string `validate:"min=8,max=100" json:"password"`
}

type LoginUserInput struct {
	UniqueId string `validate:"required" json:"uniqueId"`
	Password string `validate:"required" json:"password"`
}

type AdditionalHeader struct {
	DeviceId   string `validate:"required" json:"deviceId"`
	DeviceInfo string `validate:"required" json:"deviceInfo"`
}

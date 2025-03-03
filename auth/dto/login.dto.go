package dto

type LoginRequestDto struct {
	Email    string `json:"email" validation:"required,email"`
	Password string `json:"password" validation:"required,gte=8"`
}

type LoginResponseDto struct {
	Token string `json:"token"`
}

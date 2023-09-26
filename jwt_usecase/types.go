package jwt_usecase

type GenerateTokenInput struct {
	UserID      int
	FullName    string
	PhoneNumber int64
}

type ValidateTokenInput struct {
	Token string
}

type ValidateTokenOutput struct {
	UserID      int
	FullName    string
	PhoneNumber string
}

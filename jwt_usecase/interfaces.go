package jwt_usecase

type JWTUsecaseInterface interface {
	GenerateUserToken(input GenerateTokenInput) (string, error)
	ValidateToken(input ValidateTokenInput) (ValidateTokenOutput, error)
}

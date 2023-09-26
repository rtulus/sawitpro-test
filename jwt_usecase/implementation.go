package jwt_usecase

import (
	"fmt"

	jwt "github.com/golang-jwt/jwt/v5"
)

func (j *JWT) GenerateUserToken(input GenerateTokenInput) (string, error) {
	var token string
	key, err := jwt.ParseRSAPrivateKeyFromPEM(j.privateKey)
	if err != nil {
		return token, err
	}

	claims := make(jwt.MapClaims)
	claims["user_id"] = input.UserID
	claims["full_name"] = input.FullName
	claims["phone_number"] = fmt.Sprintf("+%d", input.PhoneNumber)

	token, err = jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	return token, err
}

func (j *JWT) ValidateToken(input ValidateTokenInput) (ValidateTokenOutput, error) {

	var output ValidateTokenOutput
	key, err := jwt.ParseRSAPublicKeyFromPEM(j.publicKey)
	if err != nil {
		return output, err
	}

	tok, err := jwt.Parse(input.Token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return key, nil
	})

	if err != nil {
		return output, fmt.Errorf("validate: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return output, fmt.Errorf("validate: invalid")
	}

	output.UserID = int(claims["user_id"].(float64))
	output.FullName = claims["full_name"].(string)
	output.PhoneNumber = claims["phone_number"].(string)

	return output, nil
}

package jwt_usecase

import "os"

type JWT struct {
	publicKey  []byte
	privateKey []byte
}

func NewJWT() *JWT {
	prvKey, err := os.ReadFile("/root/.ssh/rsakey.pem")
	if err != nil {
		panic(err)
	}
	pubKey, err := os.ReadFile("/root/.ssh/rsapubkey.pem")
	if err != nil {
		panic(err)
	}

	var jwt JWT
	jwt.privateKey = prvKey
	jwt.publicKey = pubKey

	return &jwt
}

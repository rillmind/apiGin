package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtService struct {
	secretKey string
	issure    string
}

type Claim struct {
	Sum  uint `json:"sum"`
	Role string
	jwt.StandardClaims
}

func NewJwtService() *jwtService {
	return &jwtService{
		secretKey: "SecreteKey",
		issure:    "apiGin",
	}
}

func (js *jwtService) GenerateToken(id uint) (string, error) {
	claim := &Claim{
		id,
		"user",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    js.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	token, err := to.SignedString([]byte(js.secretKey))

	if err != nil {
		return "", nil
	}

	return token, nil
}

func (js *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(to *jwt.Token) (interface{}, error) {
		if _, isValid := to.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(js.secretKey), nil
	})

	return err == nil
}

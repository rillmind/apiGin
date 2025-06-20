package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service struct {
	secretKey string
	issure    string
}

type Claim struct {
	Sum                int    `json:"sum"`
	Role               string `json:"role"`
	jwt.StandardClaims        // Embedder diretamente, sem nome de campo
}

func NewService() *Service {
	return &Service{
		secretKey: "SecreteKey",
		issure:    "apiGin",
	}
}

func (js *Service) GenerateToken(id int) (string, error) {
	claim := &Claim{
		Sum:  id,
		Role: "user",
	}
	// Como StandardClaims está embedded, podemos acessar seus campos diretamente
	claim.ExpiresAt = time.Now().Add(time.Hour * 2).Unix()
	claim.Issuer = js.issure
	claim.IssuedAt = time.Now().Unix()

	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	token, err := to.SignedString([]byte(js.secretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (js *Service) ValidateToken(tokenString string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", tokenString)
		}
		return []byte(js.secretKey), nil
	})
	return err == nil && token.Valid
}

// feito pelo copilot para eu estudar depois
func (js *Service) GetRoleFromToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(js.secretKey), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*Claim); ok && token.Valid {
		return claims.Role, nil
	}

	return "", fmt.Errorf("invalid token or claims")
}

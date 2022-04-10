package helpers

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type IJWT interface {
	GenerateToken(email string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

type authCustomClaims struct {
	Name   string `json:"name"`
	IsUser bool   `json:"isUser"`
	jwt.StandardClaims
}

type JWT struct {
	secret string
	issure string
}

func JWTAuthService() IJWT {
	return &JWT{
		secret: getSecretKey(),
		issure: "http://localhost:8080",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (j *JWT) GenerateToken(email string, isUser bool) string {
	claims := &authCustomClaims{
		Name:   email,
		IsUser: isUser,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    j.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(j.secret))
	return tokenString
}

func (j *JWT) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &authCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
}

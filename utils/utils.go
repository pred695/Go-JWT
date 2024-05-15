package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pred695/Go-JWT/Models"
	"golang.org/x/crypto/bcrypt"
)

type (
	CustomClaims struct {
		Username string
		UserId   uint

		jwt.RegisteredClaims
	}
)

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error in hashing", err)
		return ""
	}
	return string(hash)
}

func GenerateToken(user *Models.User) (string, error) {

	claims := CustomClaims{
		user.Username,
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Minute * 10)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		fmt.Println("Error in signing token", err)
		return "", err
	}

	return signedToken, nil

}

func ValidateToken(tokenString string) (*CustomClaims, error) {
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		fmt.Println("Error in validating token", err)
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)

	if !ok {
		fmt.Println("Error in validating token", err)

		return nil, err
	}

	return claims, nil
}

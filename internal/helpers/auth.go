package helpers

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CreateToken(userId string, secretKey string) (string, error) {
	claims := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": userId,
			"exp": time.Now().Add(time.Hour).Unix(),
			"iat": time.Now().Unix(),
		},
	)

	token, err := claims.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyToken(tokenString string, secretKey string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	return token, nil
}

func ParseToken(token *jwt.Token, secretKey string) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return "", errors.New("Invalid token")
	}

	sub, ok := claims["sub"].(string)

	if !ok {
		return "", errors.New("Invalid token")
	}

	return sub, nil
}

func GeneratePassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashedPassword), err
}

func VerifyPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

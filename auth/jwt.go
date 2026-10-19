package auth

import (
	"Assignment4/model"
	"errors"
	// "go/token"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("ganti-env")

type Claims struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(u *model.User) (string, error) {
	claims := Claims{
		ID: u.ID,
		Name: u.Name,
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1*time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ParseToken(tokenString string) (*Claims, error){
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(t *jwt.Token) (interface{}, error){return secret, nil},
		jwt.WithValidMethods([]string{"HS256"}), // menolak algoritma lain
	)
	if err != nil{
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid{
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
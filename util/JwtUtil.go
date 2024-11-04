package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const exp = time.Hour * 24

var jwtKey = []byte("StarFall")

func GenerateToken(user string) (token string) {
	claims := jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(exp).Unix(),
	}
	claimWithSign := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := claimWithSign.SignedString(jwtKey)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return token
}

func ParseToken(token string) (claim jwt.MapClaims, err error) {
	_, err = jwt.ParseWithClaims(token, &claim, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return claim, err
}

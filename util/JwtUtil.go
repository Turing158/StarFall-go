package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserClaim struct {
	User  string
	Email string
	Role  string
}

const exp = time.Hour * 24

var jwtKey = []byte("StarFall")

func GenerateToken(claim UserClaim) (token string) {
	claims := jwt.MapClaims{
		"user":  claim.User,
		"email": claim.Email,
		"role":  claim.Role,
		"exp":   time.Now().Add(exp).Unix(),
	}
	return Generate(claims)
}

func GenerateTokenWithExpire(claim UserClaim, expire time.Duration) (token string) {
	claims := jwt.MapClaims{
		"user":  claim.User,
		"email": claim.Email,
		"role":  claim.Role,
		"exp":   time.Now().Add(expire).Unix(),
	}
	return Generate(claims)
}

func Generate(claims jwt.MapClaims) (token string) {
	claimWithSign := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := claimWithSign.SignedString(jwtKey)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return token
}

func ParseToken(token string) (claim jwt.MapClaims, userClaim UserClaim, err error) {
	_, err = jwt.ParseWithClaims(token, &claim, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	userClaim.User = claim["user"].(string)
	userClaim.Email = claim["email"].(string)
	userClaim.Role = claim["role"].(string)
	return
}

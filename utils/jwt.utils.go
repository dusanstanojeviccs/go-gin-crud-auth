package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

type AppClaims struct {
	UserId int `json:"userId"`
	jwt.RegisteredClaims
}

type jwtUtils struct {
}

var key = []byte("THIS_WILL_BE_OUR_ENV_VAR_SOON")

func signer(t *jwt.Token) (interface{}, error) {
	return key, nil
}

func (this *jwtUtils) GenerateSessionCookie(userId int) string {
	value, error := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
	}).SignedString(key)

	if error != nil {
		return ""
	}

	return value
}

func (this *jwtUtils) ParseSessionCookie(auth string) (userId int, valid bool) {
	valid = true

	token, error := jwt.ParseWithClaims(auth, &AppClaims{}, signer)

	if error != nil {
		valid = false
		return
	}

	if !token.Valid {
		valid = false
		return
	}

	if claims, ok := token.Claims.(*AppClaims); ok {
		valid = true
		userId = claims.UserId
	}

	return
}

var Jwt = jwtUtils{}

package jwt

import (
	"go-gin-crud-auth/utils"

	"github.com/golang-jwt/jwt/v5"
)

type AppClaims struct {
	UserId int `json:"userId"`
	jwt.RegisteredClaims
}

type jwtUtils struct {
	key []byte
}

func (this *jwtUtils) signer(t *jwt.Token) (interface{}, error) {
	return this.key, nil
}

func (this *jwtUtils) GenerateSessionCookie(userId int) string {
	value, error := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
	}).SignedString(this.key)

	if error != nil {
		return ""
	}

	return value
}

func (this *jwtUtils) ParseSessionCookie(auth string) (userId int, valid bool) {
	valid = true

	token, error := jwt.ParseWithClaims(auth, &AppClaims{}, this.signer)

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

var Jwt *jwtUtils

func Init() {
	if Jwt != nil {
		panic("JWT has already been initialized")
	}
	Jwt = &jwtUtils{key: []byte(utils.Config.Server.JwtSecret)}
}

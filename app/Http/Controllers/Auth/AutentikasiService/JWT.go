package AutentikasiService

import (
	"login-sistem-jwt/app/Http/Controllers/User/UserModel"
	"login-sistem-jwt/app/Provider/ErrorHandler"
	"time"

	"github.com/golang-jwt/jwt"
)

func GetJWT(dataUser UserModel.User) map[string]interface{} {
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	claims := &jwt.StandardClaims{
		Audience:  dataUser.Name,
		ExpiresAt: expiresAt,
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	tokenString, err := token.SignedString([]byte("secret"))
	ErrorHandler.Err(err).Check("SignerString Gagal").Error()

	response := make(map[string]interface{})

	response["token"] = tokenString
	response["user"] = dataUser

	return response
}
package AutentikasiService

import (
	"fmt"
	controllers "inventori/app/Http/Controllers"
	"inventori/app/Provider/ResponseHandler"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func GetJWT(dataUser map[string]interface{}) (map[string]interface{}, bool) {
	response := make(map[string]interface{})

	expiredTime, _ := strconv.Atoi(os.Getenv("TokenExpired"))
	fmt.Println(expiredTime)
	expiresAt := time.Now().Add(time.Minute * time.Duration(expiredTime)).Unix()

	nameOftheUser := dataUser["name"]
	claims := &jwt.StandardClaims{
		Audience:  nameOftheUser.(string),
		ExpiresAt: expiresAt,
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		ResponseHandler.Go(controllers.GlobalGContext).CustomProccessFailure(err.Error())
		return response, false
	}

	response["token"] = tokenString
	response["user"] = dataUser

	return response, true
}

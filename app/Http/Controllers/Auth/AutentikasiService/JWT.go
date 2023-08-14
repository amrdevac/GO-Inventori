package AutentikasiService

import (
	controllers "inventori/app/Http/Controllers"
	"inventori/app/Http/Controllers/User/UserModel"
	"inventori/app/Provider/ResponseHandler"
	"time"

	"github.com/golang-jwt/jwt"
)

func GetJWT(dataUser UserModel.User) (map[string]interface{}, bool) {
	response := make(map[string]interface{})

	expiresAt := time.Now().Add(time.Minute * 1).Unix()

	claims := &jwt.StandardClaims{
		Audience:  dataUser.Name,
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

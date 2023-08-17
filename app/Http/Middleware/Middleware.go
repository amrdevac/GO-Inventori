package Middleware

import (
	"inventori/app/Provider/ResponseHandler"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func VerifyJWT(c *gin.Context) {
	header := c.Request.Header["Token"]

	if header == nil {
		ResponseHandler.Go(c).
			SetMessage("Unauthorize").
			SetHttpStatus(http.StatusBadRequest).
			Apply().StopProcess()
	}

	claims := &jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(header[0], claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		
	if err != nil {
		ResponseHandler.Go(c).
			SetMessage("Access Unauthorized").
			SetHttpStatus(http.StatusUnauthorized).
			Apply().StopProcess()
	}
}

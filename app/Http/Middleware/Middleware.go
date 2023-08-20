package Middleware

import (
	"inventori/app/Provider/ResponseHandler"
	"inventori/pkg/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func VerifyJWT(c *gin.Context) {
	header := c.Request.Header["Authorization"]
	if !helper.CompareChecker(strings.Join(header, ""), "==", "", "Authorization required") {
		c.Abort()
	}

	hasil := strings.TrimSpace(strings.Replace(header[0], "Bearer", "", -1))

	if header == nil {
		ResponseHandler.Go(c).
			SetMessage("Unauthorize").
			SetHttpStatus(http.StatusBadRequest).
			Apply()
		c.Abort()

	}

	claims := &jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(hasil, claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

	if err != nil {
		ResponseHandler.Go(c).
			SetMessage("Access Unauthorized").
			SetHttpStatus(http.StatusUnauthorized).
			Apply()
		c.Abort()

	}

	c.Next()

}

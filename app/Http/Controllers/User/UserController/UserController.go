package UserController

import (
	"login-sistem-jwt/app/Http/Controllers/User/UserModel"
	"login-sistem-jwt/app/Provider/RequestJson"
	"login-sistem-jwt/app/Provider/ResponseHandler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	users := []UserModel.User{}
	UserModel.GormConnect.Find(&users)
	ResponseHandler.Go(c).SetData(users).SetHttpStatus(http.StatusOK).Get()
}

func Show(c *gin.Context) {
	user := UserModel.User{}
	email := c.Param("email")
	UserModel.GormConnect.Where("email = ?", email).Take(&user)
	ResponseHandler.Go(c).SetData(user).SetHttpStatus(http.StatusOK).Get()
}

func Update(c *gin.Context) {
	userModel, userRequest := UserModel.User{}, UserModel.UserUpdate{}

	err := c.ShouldBindJSON(&userRequest)
	RequestJson.Validate(err, c)

	email := c.Param("email")
	UserModel.GormConnect.Where("email = ?", email).First(&userModel).Updates(userRequest)
	ResponseHandler.Go(c).SetData(userModel).SetHttpStatus(http.StatusOK).Get()
}

func Destroy(c *gin.Context) {
	userModel := UserModel.User{}
	email := c.Param("email")

	UserModel.GormConnect.Where("email = ?", email).Delete(&userModel)
	ResponseHandler.Go(c).SetHttpStatus(http.StatusOK).Get()
}

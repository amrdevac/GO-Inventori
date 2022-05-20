package UserModel

import (
	"login-sistem-jwt/app/Database/Mysql"
)

var GormConnect = Mysql.Connect()

type User struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Password string `json:"password" binding:"required"`
}

type UserUpdate struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
package UserModel

import (
	"inventori/app/Database/Mysql"
)

var GormConnect = Mysql.Connect()

type User struct {
	UserID   string `json:"user_id" gorm:"primaryKey;autoIncrement:true"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Password string `json:"password" binding:"required"`
}

type UserUpdate struct {
	Name   string `json:"name" binding:"required"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

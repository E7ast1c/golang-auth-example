package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"type:varchar(100);unique_index"`
	Gender   string `json:"Gender"`
	Password string `json:"Password"`
}

type Token struct {
	UserID uint
	Name   string
	Email  string
	*jwt.StandardClaims
}

type Exception struct {
	Message string `json:"message"`
}

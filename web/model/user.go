package model

import (
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email" validate:"required"`
	Password string `json:"password"`
	IsActive bool `json:"is_active"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type UserResponse struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
}

type UpdateInfo struct {
    Username string `json:"username"`
    Password   string `json:"email"`
}

type OTPRequest struct {
	Email    string `json:"email" validate:"required"`
	OTP  string `json:"otp" validate:"required"`
}

func (u *User)HashPassword() error{
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err 
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err 
	}
	return nil 
}



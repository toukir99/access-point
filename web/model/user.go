package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	Username string `json:"username" validate:"required"`
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (u *User)HashPassword() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)
}
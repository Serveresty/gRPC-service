package service

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Login        string
	HashPassword string
}

func NewUser(login string, password string) (*User, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		Login:        login,
		HashPassword: string(hashPass),
	}

	return user, nil
}

func (user *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(password))
	return err == nil
}

func (user *User) Clone() *User {
	return &User{
		Login:        user.Login,
		HashPassword: user.HashPassword,
	}
}

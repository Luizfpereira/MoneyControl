package entity

import (
	"errors"
	"fmt"
	"net/mail"
	"unicode"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(name, lastName, email, password string) (*User, error) {
	user := &User{
		Name:     name,
		LastName: lastName,
		Email:    email,
		Password: password,
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("name is empty")
	}
	if u.LastName == "" {
		return errors.New("last name is empty")
	}
	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return fmt.Errorf("einvalid email: %v", err)
	}

	if !isPasswordValid(u.Password) {
		return errors.New("invalid password")
	}
	return nil
}

func isPasswordValid(p string) bool {
	var (
		hasNumber  = false
		hasUpper   = false
		hasLower   = false
		hasSpecial = false
	)
	if len(p) < 7 {
		return false
	}

	for _, c := range p {
		switch {
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}
	return hasNumber && hasUpper && hasLower && hasSpecial
}

package entity

import (
	"errors"
	"fmt"
	"net/mail"
	"unicode"
)

// User defines an object of a user who will be responsible for the transactions in the application
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// NewUser returns a User according to the name, last name, email and password passed to the function
// or an error if invalid arguments are passed
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

// Validate verifies if the values used to create a User object are valid
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

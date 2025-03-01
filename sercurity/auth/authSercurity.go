package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var users = map[string]string{}

// register nguoi dung moi
func Register(username, password string) error {
	if _, exists := users[username]; exists {
		return errors.New("username already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	users[username] = string((hashedPassword))
	return nil
}

// Authentication
func Authentication(username, password string) bool {
	hashedPassword, exists := users[username]
	if !exists {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

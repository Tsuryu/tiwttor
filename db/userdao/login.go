package userdao

import (
	"github.com/Tsuryu/tiwttor/model"
	"golang.org/x/crypto/bcrypt"
)

// Login : logs an user
func Login(email string, password string) (model.User, bool) {
	user, result, _ := FindBy(email)
	if !result {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}

	return user, result
}
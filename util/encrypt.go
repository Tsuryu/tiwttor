package util

import "golang.org/x/crypto/bcrypt"

// Encrypt : use dcrypt to encrypt data
func Encrypt(value string) (string, error) {
	cost := 8 // amount of times that the value will be encrypted, 2 to 8 = 256, 6 = normal user, 8 = admin user
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), cost)
	return string(bytes), err
}

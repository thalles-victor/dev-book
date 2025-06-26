package security

import "golang.org/x/crypto/bcrypt"

// hsh recive a string an pu a hash
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 12)
}

// Compare if password matches
func CheckHashPassword(password, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

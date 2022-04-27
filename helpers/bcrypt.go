package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(p string) string {
	salt := 8
	password := []byte(p)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hashedPassword)
}

func ComparePassword(h, p []byte) bool {
	hash, password := []byte(h), []byte(p)
	err := bcrypt.CompareHashAndPassword(hash, password)

	return err == nil
}

package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(p string) string {
	salt := bcrypt.DefaultCost
	password := []byte(p)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hashedPassword)
}

func ComparePassword(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}

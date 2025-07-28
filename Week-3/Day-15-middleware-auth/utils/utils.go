package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"github.com/nazartymiv/go-mastery/Week-3/Day-15-middleware/logger"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		logger.Error("Failed to generate token", err.Error())
		log.Fatal()
	}
	return base64.URLEncoding.EncodeToString(bytes)
}

package hash

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

const cost = 10

func HashPassword(password string) (string, error) {
	log.Printf("Hashing password: %s", password)
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	hash := string(hashedBytes)
	log.Printf("Generated hash: %s", hash)
	return hash, nil
}

func CheckPassword(password, hashedPassword string) bool {
	log.Printf("CheckPassword - password bytes: %v", []byte(password))
	log.Printf("CheckPassword - hashedPassword bytes: %v", []byte(hashedPassword))
	
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Printf("Password check failed: %v", err)
		return false
	}
	return true
}

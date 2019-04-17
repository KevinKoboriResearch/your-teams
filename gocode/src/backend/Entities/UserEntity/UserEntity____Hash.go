package UserEntity

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

//__ GENERATE HASH PASSWORD __________________________________________________//
func GenerateHashPassword(password string) (string, error) {
/*________________________________________TESTING FUNCTION________________________________________*/
log.Println("\n\nUserEntity.GenerateHashPassword\n")

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 1)

	return string(bytes), err
}

//__ CHECK HASH PASSWORD _____________________________________________________//
func CheckPasswordHash(password string, hash string) bool {
/*________________________________________TESTING FUNCTION________________________________________*/
log.Println("\n\nUserEntity.CheckPasswordHash\n")

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

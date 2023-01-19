package services

import (
	"errors"
	"net/http"
	"serviceOauth/models"
	"strconv"

	"serviceOauth/auth"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func hashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err

}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUser(db *gorm.DB, id uint) (models.User, error) {
	var user models.User
	result := db.Where("id = ?", id).First(&user)
	return user, result.Error
}

func RegisterUser(db *gorm.DB, user models.User) (models.User, error) {

	hashedPassword, _ := hashPassword(user.Password)
	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
	}
	result := db.Create(&newUser)
	return newUser, result.Error
}

func LoginUser(db *gorm.DB, email string, password string) (models.User, error) {

	var user models.User
	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	if !checkPasswordHash(password, user.Password) {
		return user, errors.New("Invalid Password")
	}

	token, err := auth.CreateToken(user.ID, user.Name, user.Email)

	if err != nil {
		return user, err
	}
	_ = db.Model(&user).Update("token", token)
	user.Token = token
	return user, nil
}

func RefreshToken(db *gorm.DB, r *http.Request) (string, error) {
	tokenExpired := auth.ExtractToken(r)
	tokenExpiredId, err := auth.ExtractTokenExpiredId(r)

	if err != nil {
		return "", err
	}
	if tokenExpiredId == 0 {
		return "", errors.New("Invalid Token")
	}

	var user models.User
	result := db.Where("id = ?", tokenExpiredId).First(&user)

	if result.Error != nil {
		return "", result.Error
	}

	if tokenExpired != user.Token {
		return "", errors.New("Invalid Token")
	}

	convString := strconv.Itoa(int(tokenExpiredId))

	newToken, err := auth.RefreshToken(convString, tokenExpiredId)

	if err != nil {
		return "", err
	}

	_ = db.Model(&user).Update("token", newToken)
	return newToken, nil

}

func VerifyAuth(db *gorm.DB, r *http.Request) error {
	err := auth.TokenValid(r)
	if err != nil {
		return err
	}
	return nil
}

func ExtractTokenId(r *http.Request) (uint, error) {
	tokenId, err := auth.ExtractTokenId(r)
	if err != nil {
		return 0, err
	}
	return tokenId, nil
}

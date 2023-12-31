package repository

import (
	"errors"
	"on-air/models"
	"on-air/utils"

	"gorm.io/gorm"
)

func GetUserByEmail(db *gorm.DB, email string) (*models.User, error) {
	var dbUser models.User
	result := db.First(&dbUser, "email = ?", email)
	if result.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	return &dbUser, nil
}

func RegisterUser(db *gorm.DB, email string, password string) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := models.User{Email: email, Password: hashedPassword}
	result := db.Create(&user)
	if err = result.Error; err != nil {
		return nil, err
	}

	return &user, nil
}

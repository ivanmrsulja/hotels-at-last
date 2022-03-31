package repository

import (
	"errors"
	"strconv"
	"time"

	model "github.com/ivanmrsulja/hotels-at-last/user-service/model"
	utils "github.com/ivanmrsulja/hotels-at-last/user-service/utils"
)

func FindUserById(id uint) (model.User, error) {
	var user model.User

	utils.Db.First(&user, id)

	if user.ID == 0 {
		return user, errors.New("There is no user with ID " + strconv.FormatUint(uint64(id), 10))
	}

	return user, nil
}

func BanUser(id uint, endDate time.Time) error {
	var user model.User

	utils.Db.First(&user, id)

	if user.ID == 0 {
		return errors.New("There is no user with ID " + strconv.FormatUint(uint64(id), 10))
	}

	user.BannedUntil = endDate
	utils.Db.Save(&user)

	return nil
}

func CreateUser(user model.User) (model.User, error) {

	user.Role = model.USER
	user.BannedUntil = time.Now()
	createdUser := utils.Db.Create(&user)

	if createdUser.Error != nil {
		return user, createdUser.Error
	}

	return user, nil
}

func CheckCredentials(email string, password string) (model.User, error) {
	var user model.User

	utils.Db.Table("users").Where("email = ? AND password = ?", email, password).First(&user)

	if user.ID == 0 {
		return user, errors.New("Invalid username or password.")
	}

	if time.Now().Before(user.BannedUntil)  {
		return user, errors.New("You are banned until: " + user.BannedUntil.String())
	}

	return user, nil
}
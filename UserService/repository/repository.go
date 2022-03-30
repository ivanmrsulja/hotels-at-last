package repository

import (
	"errors"
	"strconv"

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

func BanUser(id uint) {

}

func CheckCredentials(email string, password string) (model.User, error) {
	var user model.User

	utils.Db.Table("users").Where("email = ? AND password = ?", email, password).First(&user)

	if user.ID == 0 {
		return user, errors.New("Invalid username or password.")
	}

	return user, nil
}
package database

import (
	"asidikfauzi/reservation-of-sport-fields-golang/config"
	"asidikfauzi/reservation-of-sport-fields-golang/models"
)

func InsertUsers(users models.Users) (interface{}, error) {
	if err := config.DB.Save(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

package repository

import (
	"fmt"
	"log"
	"resto-be/database"
	"resto-be/database/dbmodels"
	"resto-be/models/dto"
)

// GetUserByEmail ...
func GetUserByEmail(email string) (dbmodels.User, error) {
	db := database.GetDbCon()
	db.Debug().LogMode(true)

	var user dbmodels.User
	var err error

	db.Where("email = ?", email).Find(&user)

	fmt.Println("User => ", user)
	return user, err

}

// GetUserFilterPaging ...
func GetUserFilterPaging(req dto.UserRequesDto, page int, limit int) ([]dbmodels.User, int, error) {
	db := database.GetDbCon()

	var users []dbmodels.User
	var total int

	if err := db.Limit(limit).Offset((page - 1) * limit).Order("id").Find(&users).Limit(-1).Offset(0).Count(&total).Error; err != nil {
		log.Println("<<< Error get data user by filter paging >>>")

		return users, 0, err
	}

	for _, user := range users {
		user.Password = "****"
	}
	fmt.Println("iterate user ", users)

	return users, total, nil
}

func SaveUser(user *dbmodels.User) error {
	db := database.GetDbCon()

	err := db.Save(&user).Error

	return err
}

package repository

import (
	"fmt"
	"resto-be/database"
	"resto-be/database/dbmodels"
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
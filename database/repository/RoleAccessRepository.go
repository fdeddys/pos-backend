package repository

import (
	"log"
	"resto-be/database"
	"strconv"
)

func CountUserAccessByUserIdAndAccessName(userId int64, accessName string) bool {
	db := database.GetDbCon()
	var total int
	user := "user"

	err := db.Raw("SELECT COUNT (*) FROM role_access JOIN ROLE ON role_access.role_id=ROLE.ID JOIN ACCESS ON role_access.access_id=ACCESS.ID  JOIN " + strconv.Quote(user) + " ON ROLE.ID= " + strconv.Quote(user) + ".role_id WHERE " + strconv.Quote(user) + ".ID=? AND ACCESS.NAME=?", userId, accessName).Count(&total).Error

	if err != nil {
		log.Println("err ====> ", err)
		return false
	}
	if total == 0 {
		return false
	}
	return true
}

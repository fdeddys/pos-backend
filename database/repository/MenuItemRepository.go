package repository

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"resto-be/database"
	"resto-be/database/dbmodels"
	"resto-be/models/dto"
)

func SaveMenuItem(menuItem *dbmodels.MenuItem) (error) {
	db := database.GetDbCon()

	err := db.Save(&menuItem).Error

	return err
}

func SaveMenuItemPicture(picture *dbmodels.MenuItemPicture) error  {
	db := database.GetDbCon()

	err := db.Save(&picture).Error

	return err
}

func GetAllMenuItem() ([]dbmodels.MenuItem, error) {
	db := database.GetDbCon()

	var menuItems []dbmodels.MenuItem

	err := db.Find(&menuItems).Error

	return menuItems, err
}

func GetMenuItemById(id int64) (dbmodels.MenuItem, error) {
	db := database.GetDbCon()

	var menuItem dbmodels.MenuItem

	if id == 0 {
		return menuItem, errors.New("id = 0")
	}
	err := db.Where(dbmodels.MenuItem{ID:id}).First(&menuItem).Error
	return menuItem, err
}

func GetMenuItemByMenuGroupId(id int64) ([]dbmodels.MenuItem, error) {
	db := database.GetDbCon()

	var menuItems []dbmodels.MenuItem

	if id == 0 {
		return menuItems, errors.New("id = 0")
	}
	err := db.Where(dbmodels.MenuItem{GroupID:id}).Find(&menuItems).Error
	return menuItems, err
}

func GetMenuItemFilterPaging(req dto.MenuItemRequestDto, page int, limit int) ([]dbmodels.MenuItem, int, error) {
	db := database.GetDbCon()

	var menuItems []dbmodels.MenuItem

	var total int
	var whereQuery string

	query := " FROM e_menu_item " +
		"JOIN e_menu_group ON e_menu_group.ID=e_menu_item.group_id " +
		"JOIN resto ON e_menu_group.resto_id=resto.ID "

	if req.RestoId != 0 {
		whereQuery = fmt.Sprintf("and resto.id = %v", req.RestoId)

	}

	querySelect := "select * " + query + whereQuery
	count := "select count(*) " + query + whereQuery

	errSelectChan := make(chan error)
	countChan := make(chan int)
	errCountChan := make(chan error)

	go AsyncSelectQueryMenuItem(db, limit, page, querySelect, &menuItems, errSelectChan)
	go AsyncCountQueryMenuItem(db, count, total, errCountChan, countChan)

	err :=<- errSelectChan
	errCount :=<- errCountChan
	total = <-countChan

	//errCount := db.Raw(count).Count(&total).Error

	//err := db.Limit(limit).Offset((page-1) * limit).Where(whereQuery).Order("id").Find(&menuItems).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		log.Println("<<< Error get data menu item by filter paging >>>")
		return menuItems, 0, err
	}

	if errCount != nil {
		log.Println("<<< Error count data menu item by filter paging >>>")
		return menuItems, 0, errCount
	}


	return menuItems, total, err
}

func AsyncSelectQueryMenuItem(db *gorm.DB, limit int, page int, querySelect string, menuItems *[]dbmodels.MenuItem,  err chan error)  {
	err <- db.Limit(limit).Offset((page-1) * limit).Raw(querySelect).Scan(&menuItems).Error
	close(err)
}

func AsyncCountQueryMenuItem(db *gorm.DB, query string, total int,  err chan error, count chan int)  {
	err <- db.Raw(query).Count(&total).Error
	count <- total
	close(err)
	close(count)
}

/*

func GetMenuItemByMenuGroupIdAndRestoId(groupId int64, restoId int64) ([]dbmodels.MenuItem, error) {
	db := database.GetDbCon()

	var menuItems []dbmodels.MenuItem

	if groupId == 0 || restoId == 0{
		return menuItems, errors.New("id = 0")
	}
	err := db.Where(dbmodels.MenuItem{GroupID:groupId, RestoID:restoId}).Find(&menuItems).Error
	return menuItems, err
}
*/
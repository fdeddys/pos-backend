package dbmodels

type MenuItem struct {
	ID      int64   `json:"id"`
	GroupID int64   `json:"groupId"`
	MenuGroup 	MenuGroup `json:"menuGroup" gorm:"foreignkey:id; association_foreignkey:GroupId; association_autoupdate:false;association_autocreate:false"`
	RestoID int64 	`json:"restoId"`
	Name    string  `json:"name"`
	Desc    string  `json:"desc"`
	ImgUrl  string  `json:"imgURL"`
	Price   float64 `json:"price"`
	Stock   int `json:"stock"`
	Status  int     `json:"status"`
	IsFavorite  int     `json:"isFavorite"`
	CategoryId  int64 `json:"categoryId" gorm:"category_id"`
	Category 	Category `json:"category" gorm:"foreignkey:id; association_foreignkey:CategoryId; association_autoupdate:false;association_autocreate:false"`
	Pictures []MenuItemPicture `json:"pictures,omitempty"gorm:"foreignkey:menuItemId"`


}

// TableName ...
func (t *MenuItem) TableName() string {
	return "public.e_menu_item"
}


type MenuItem2 struct {
	ID      int64   `json:"id"`
	GroupID int64   `json:"groupId"`
	Name    string  `json:"name"`
	Desc    string  `json:"desc"`
	ImgUrl  string  `json:"imgURL"`
	Price   float64 `json:"price"`
	Stock   int `json:"stock"`
	Status  int     `json:"status"`
}
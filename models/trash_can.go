package models

type TrashCan struct {
	ID       uint   `json:id`
	Name     string `json:name`
	Location string `json:location`
	UserID   string `json:userid`
	User     User   `json:users;gorm:"foreignKey:UserID"`
}

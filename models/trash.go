package models

type Trash struct {
	ID         uint     `json:id`
	Classified string   `json:classified`
	TrashCanID uint     `json:trashcanid`
	Image      string   `json:image`
	TrashCan   TrashCan `json:trashcan";gorm:"foreignKey:TrashCanID"`
}

package models

type Trash struct {
	ID         uint     `json:"id" form:"id"`
	Classified string   `json:"classified" form:"classified"`
	TrashCanID uint     `json:"trashcanid" form:"trashcanid"`
	Image      string   `json:"image" form:"image"`
	TrashCan   TrashCan `json:"trashcan";gorm:"foreignKey:TrashCanID"`
}

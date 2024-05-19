package models

type Grouping struct {
	ID          int64    `gorm:"primaryKey;column:ID"`
	ObjectValue string   `json:"objectValue" gorm:"column:ObjectValue"`
	Pepole      []Pepole `gorm:"foreignKey:PepoleGroupingID;references:ID"`
}

package models

type Grouping struct {
	ID          int64    `json:"id" gorm:"primaryKey;column:ID"`
	ObjectValue string   `json:"objectValue" gorm:"column:ObjectValue"`
	Pepoles     []Pepole `json:"pepoles" gorm:"foreignKey:ID"`
}

func (Grouping) TableName() string {
	return "Grouping"
}

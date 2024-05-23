package models

type Grouping struct {
	ID          int64    `gorm:"primaryKey;column:ID"`
	ObjectValue string   `json:"objectValue" gorm:"column:ObjectValue"`
	Pepoles     []Pepole `gorm:"foreignKey:ID"`
}

func (Grouping) TableName() string {
	return "Grouping"
}

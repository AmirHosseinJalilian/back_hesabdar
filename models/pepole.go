package models

type Pepole struct {
	ID                 int64               `json:"id" gorm:"primaryKey;column:ID"`
	Name               string              `json:"name" gorm:"column:Name"`
	PepoleType         string              `json:"pepoleType" gorm:"column:PepoleType"`
	CodPepole          string              `json:"codPepole" gorm:"column:CodPepole"`
	GroupingID         uint                `json:"groupingID" gorm:"column:ID"`
	PepoleDescriptions []PepoleDescription `json:"pepoleDescriptions" gorm:"foreignKey:ID"`
}

func (Pepole) TableName() string {
	return "Pepole"
}

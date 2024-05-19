package models

type PepoleDescription struct {
	ID              int64  `gorm:"primaryKey;column:ID"`
	PepoleID        int64  `json:"pepoleID" gorm:"column:PepoleID"`
	Address         string `json:"address" gorm:"column:Address"`
	Phone           string `json:"phone" gorm:"column:Phone"`
	NationalityCode string `json:"nationalityCode" gorm:"column:NationalityCode"`
}

package models

type Commoditym struct {
	ID            int64  `json:"id" gorm:"primaryKey;column:ID"`
	ComodityCod   string `json:"comodityCod" gorm:"column:ComodityCod"`
	CommodityName string `json:"commodityName" gorm:"column:CommodityName"`
	UnitCount     int64  `json:"unitCount" gorm:"column:UnitCount"`
	BasePrice     int64  `json:"basePrice" gorm:"column:BasePrice"`
}

func (Commoditym) TableName() string {
	return "Commoditym"
}

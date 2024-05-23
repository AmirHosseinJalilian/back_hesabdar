package models

type SaleFactorConfirmationDetails struct {
	ID                       int64      `json:"id" gorm:"primaryKey;column:ID"`
	SaleFactorConfirmationID int64      `json:"saleFactorConfirmationID" gorm:"column:SaleFactorConfirmationID"`
	Count                    string     `json:"count" gorm:"column:Count"`
	UnitCost                 float64    `json:"unitCost" gorm:"column:UnitCost"`
	CommodityDiscount        string     `json:"commodityDiscount" gorm:"column:CommodityDiscount"`
	ISCommodityDiscount      string     `json:"iSCommodityDiscount" gorm:"column:ISCommodityDiscount"`
	Vat                      string     `json:"vat" gorm:"column:Vat"`
	CommodityID              int64      `json:"commodityID" gorm:"column:Commodity"`
	Commodity                Commoditym `gorm:"foreignKey:CommodityID"`
}

func (SaleFactorConfirmationDetails) TableName() string {
	return "SaleFactorConfirmationDetails"
}

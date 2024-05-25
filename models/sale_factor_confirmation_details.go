package models

type SaleFactorConfirmationDetails struct {
	ID                       int64      `json:"id" gorm:"primaryKey;column:ID"`
	SaleFactorConfirmationID int64      `json:"saleFactorConfirmationID" gorm:"column:SaleFactorConfirmationID"`
	Count                    float64    `json:"count" gorm:"column:Count"`
	UnitCost                 float64    `json:"unitCost" gorm:"column:UnitCost"`
	CommodityDiscount        float64    `json:"commodityDiscount" gorm:"column:CommodityDiscount"`
	ISCommodityDiscount      bool       `json:"iSCommodityDiscount" gorm:"column:ISCommodityDiscount"`
	Vat                      float64    `json:"vat" gorm:"column:Vat"`
	CommodityID              float64    `json:"commodityID" gorm:"column:Commodity"`
	Commodity                Commoditym `json:"commodity" gorm:"foreignKey:CommodityID"`
}

func (SaleFactorConfirmationDetails) TableName() string {
	return "SaleFactorConfirmationDetails"
}

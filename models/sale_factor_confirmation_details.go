package models

type SaleFactorConfirmationDetails struct {
	ID                       int64     `json:"id" gorm:"primaryKey;column:ID"`
	SaleFactorConfirmationID int64     `json:"saleFactorConfirmationID" gorm:"column:SaleFactorConfirmationID"`
	Count                    string    `json:"count" gorm:"column:Count"`
	UnitCost                 float64   `json:"unitCost" gorm:"column:UnitCost"`
	CommodityDiscount        string    `json:"commodityDiscount" gorm:"column:CommodityDiscount"`
	ISCommodityDiscount      string    `json:"iSCommodityDiscount" gorm:"column:ISCommodityDiscount"`
	Vat                      string    `json:"vat" gorm:"column:Vat"`
	CommodityID              uint      `json:"commodityId" gorm:"column:Commodity"`
	Commodity                Commodity `gorm:"foreignKey:CommodityID;references:ID"`
}

func (SaleFactorConfirmationDetails) TableName() string {
	return "SaleFactorConfirmationDetails"
}

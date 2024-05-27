package models

import "time"

type SaleFactorConfirmationDetails struct {
	DRowID                   string     `json:"drowId"`
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

type SaleFactorConfirmation struct {
	ID               int64                           `json:"id" gorm:"primaryKey;column:ID"`
	RowID            string                          `json:"rowId"`
	DateFactorSale   time.Time                       `json:"dateFactorSale" gorm:"column:DateFactorSale"`
	FactorNumber     string                          `json:"factorNumber" gorm:"column:FactorNumber"`
	SaleType         int                             `json:"saleType" gorm:"column:SaleType"`
	PepoleGroupingID int                             `json:"pepoleGroupingID" gorm:"column:PepoleGroupingID"`
	Details          []SaleFactorConfirmationDetails `gorm:"foreignKey:SaleFactorConfirmationID"`
	PepoleGrouping   Grouping                        `gorm:"foreignKey:ID"`
}

func (SaleFactorConfirmation) TableName() string {
	return "SaleFactorConfirmation"
}

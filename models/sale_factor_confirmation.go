package models

import (
	"time"
)

type SaleFactorConfirmation struct {
	ID               int64                           `json:"id" gorm:"primaryKey;column:ID"`
	RowID            string                          `json:"rowId"`
	DateFactorSale   time.Time                       `json:"dateFactorSale" gorm:"column:DateFactorSale"`
	FactorNumber     string                          `json:"factorNumber" gorm:"column:FactorNumber"`
	SaleType         int                             `json:"saleType" gorm:"column:SaleType"`
	PepoleGroupingID int                             `json:"pepoleGroupingID" gorm:"column:PepoleGroupingID"`
	Grouping         Grouping                        `gorm:"foreignKey:PepoleGroupingID;references:ID"`
	Details          []SaleFactorConfirmationDetails `gorm:"foreignKey:SaleFactorConfirmationID;references:ID"`
}

func (SaleFactorConfirmation) TableName() string {
	return "SaleFactorConfirmation"
}

// type SaleFactorConfirmation struct {
// 	ID             int64     `json:"id" gorm:"primaryKey;column:ID"`
// 	RowID          string    `json:"rowId"`
// 	DateFactorSale time.Time `json:"dateFactorSale" gorm:"column:DateFactorSale"`
// 	FactorNumber   string    `json:"factorNumber" gorm:"column:FactorNumber"`
// 	SaleType       int       `json:"saleType" gorm:"column:SaleType"`
// 	// PepoleGroupingID         int64     `json:"pepoleGroupingId" gorm:"column:PepoleGroupingID"`
// 	// ObjectValue              string    `json:"objectValue" gorm:"column:ObjectValue"`
// 	// Name                     string    `json:"name" gorm:"column:Name"`
// 	// NationalityCode          string    `json:"nationalityCode" gorm:"column:NationalityCode"`
// 	// SaleFactorConfirmationID int64     `json:"saleFactorConfirmationID" gorm:"column:SaleFactorConfirmationID"` // This should match your database column name
// 	// // SaleFactorCommodities    []SaleFactorCommodity `json:"saleFactorCommodities" gorm:"foreignKey:SaleFactorConfirmationID"`
// 	// Count               float64 `json:"count" gorm:"column:Count"`
// 	// UnitCost            float64 `json:"unitCost" gorm:"column:UnitCost"`
// 	// CommodityDiscount   float64 `json:"commodityDiscount" gorm:"column:CommodityDiscount"`
// 	// ISCommodityDiscount bool    `json:"iSCommodityDiscount" gorm:"column:ISCommodityDiscount"`
// 	// Vat                 float64 `json:"vat" gorm:"column:Vat"`
// 	// Phone               string  `json:"phone" gorm:"column:Phone"`
// 	// Address             string  `json:"address" gorm:"column:Address"`
// 	// PepoleType          int16   `json:"pepoleType" gorm:"column:PepoleType"`
// }

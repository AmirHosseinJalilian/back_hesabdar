package models

import (
	"time"
)

type SaleFactorConfirmation struct {
	ID               int64                         `json:"id" gorm:"primaryKey;column:ID"`
	RowID            string                        `json:"rowId"`
	DateFactorSale   time.Time                     `json:"dateFactorSale" gorm:"column:DateFactorSale"`
	FactorNumber     string                        `json:"factorNumber" gorm:"column:FactorNumber"`
	SaleType         int                           `json:"saleType" gorm:"column:SaleType"`
	PepoleGroupingID int                           `json:"pepoleGroupingID" gorm:"column:PepoleGroupingID"`
	Details          SaleFactorConfirmationDetails `gorm:"foreignKey:SaleFactorConfirmationID"`
	PepoleGrouping   Grouping                      `json:"pepoleGrouping" gorm:"foreignKey:ID"`
}

func (SaleFactorConfirmation) TableName() string {
	return "SaleFactorConfirmation"
}

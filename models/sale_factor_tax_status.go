package models

import "time"

// SaleFactorTaxStatus model
type SaleFactorTaxStatus struct {
	SaleFactorConfirmationID int64     `json:"saleFactorConfirmationID" gorm:"primaryKey;column:SaleFactorConfirmationID;autoIncrement:false"`
	Status                   uint8     `json:"status" gorm:"column:Status"`
	StatusDate               time.Time `json:"statusDate" gorm:"column:StatusDate"`
}

func (SaleFactorTaxStatus) TableName() string {
	return "SaleFactorTaxStatus"
}

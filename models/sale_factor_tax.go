package models

import "time"

// SaleFactorTax model
type SaleFactorTax struct {
	SaleFactorConfirmationID int64     `json:"saleFactorConfirmationID" gorm:"primaryKey;column:SaleFactorConfirmationID;autoIncrement:false"`
	BillType                 bool      `json:"billType" gorm:"column:BillType"`
	PostType                 uint8     `json:"postType" gorm:"column:PostType"`
	CreationDate             time.Time `json:"creationDate" gorm:"column:CreationDate"`
	SettlementMethod         uint8     `json:"settlementMethod" gorm:"column:SettlementMethod"`
	CashAmount               float64   `json:"cashAmount" gorm:"column:CashAmount"`
	LoanAmount               float64   `json:"loanAmount" gorm:"column:LoanAmount"`
}

func (SaleFactorTax) TableName() string {
	return "SaleFactorTax"
}

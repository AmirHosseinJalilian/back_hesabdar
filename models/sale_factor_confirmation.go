package models

import (
	"time"
)

type CustomSaleFactorConfirmation struct {
	ID int64 `json:"id"`
	// RowID            string                               `json:"rowId"`
	DateFactorSale      time.Time                            `json:"dateFactorSale"`
	FactorNumber        string                               `json:"factorNumber"`
	SaleType            int                                  `json:"saleType"`
	PepoleGroupingID    int                                  `json:"pepoleGroupingID"`
	Details             []CustomSaleFactorConfirmationDetail `json:"details"`
	PepoleGrouping      CustomPepoleGrouping                 `json:"pepoleGrouping"`
	SaleFactorTax       CustomSaleFactorTax                  `json:"saleFactorTax"`
	SaleFactorTaxStatus CustomSaleFactorTaxStatus            `json:"saleFactorTaxStatus"`
}

type CustomSaleFactorConfirmationDetail struct {
	// DRowID                   string          `json:"drowId"`
	ID                       int64           `json:"id"`
	SaleFactorConfirmationID int64           `json:"saleFactorConfirmationID"`
	Count                    float64         `json:"count"`
	UnitCost                 float64         `json:"unitCost"`
	CommodityDiscount        float64         `json:"commodityDiscount"`
	ISCommodityDiscount      bool            `json:"iSCommodityDiscount"`
	Vat                      float64         `json:"vat"`
	CommodityID              float64         `json:"commodityID"`
	Commodity                CustomCommodity `json:"commodity"`
}

type CustomCommodity struct {
	ID            int64  `json:"id"`
	ComodityCod   string `json:"comodityCod"`
	CommodityName string `json:"commodityName"`
	UnitCount     int64  `json:"unitCount"`
	BasePrice     int64  `json:"basePrice"`
}

type CustomPepoleGrouping struct {
	ID          int64          `json:"id"`
	ObjectValue string         `json:"objectValue"`
	Pepoles     []CustomPepole `json:"pepoles"`
}

type CustomPepole struct {
	ID                 int64                     `json:"id"`
	Name               string                    `json:"name"`
	PepoleType         uint8                     `json:"pepoleType"`
	CodPepole          string                    `json:"codPepole"`
	GroupingID         uint                      `json:"groupingID"`
	PepoleDescriptions []CustomPepoleDescription `json:"pepoleDescriptions"`
}

type CustomPepoleDescription struct {
	ID              int64  `json:"id"`
	PepoleID        int64  `json:"pepoleID"`
	Address         string `json:"address"`
	Phone           string `json:"phone"`
	NationalityCode string `json:"nationalityCode"`
}

type CustomSaleFactorTaxStatus struct {
	SaleFactorConfirmationID int64     `json:"saleFactorConfirmationID"`
	Status                   uint8     `json:"status"`
	StatusDate               time.Time `json:"statusDate"`
}

type CustomSaleFactorTax struct {
	SaleFactorConfirmationID int64     `json:"saleFactorConfirmationID"`
	BillType                 bool      `json:"billType"`
	PostType                 uint8     `json:"postType"`
	CreationDate             time.Time `json:"creationDate"`
	SettlementMethod         uint8     `json:"settlementMethod"`
	CashAmount               float64   `json:"cashAmount"`
	LoanAmount               float64   `json:"loanAmount"`
}

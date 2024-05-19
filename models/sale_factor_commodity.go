package models

type SaleFactorCommodity struct {
	ID                       uint                    `gorm:"primaryKey;column:ID"`                       // Update to match your column name
	SaleFactorConfirmationID int64                   `gorm:"primaryKey;column:SaleFactorConfirmationID"` // Update to match your column name
	SaleFactorConfirmation   *SaleFactorConfirmation `gorm:"foreignKey:SaleFactorConfirmationID"`
	Commodity                *Commodity              `gorm:"foreignKey:CommodityID"`
}

// TableName overrides the table name used by SaleFactorCommodity to `SaleFactorCommodity`
func (SaleFactorCommodity) TableName() string {
	return "SaleFactorCommodity" // Update to match your join table name
}

package commodity

type Commodity struct {
	ID              uint   `gorm:"primaryKey;uniqueIndex:udx_Commoditys"`
	ComodityCod     string `json:"comodityCod"`
	CommodityName   string `json:"commodityName"`
	UnitCount       string `json:"unitCount"`
	BasePrice       int64  `json:"basePrice"`
	CommodityNumber string `json:"commodityNumber"`
}

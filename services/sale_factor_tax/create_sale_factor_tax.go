package sale_factor_tax

import (
	"github.com/AmirHosseinJalilian/back_hesabdar/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func CreateSaleFactorTax(c echo.Context, db *gorm.DB) error {
	var requestBody models.CustomSaleFactorTax

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request body",
		})
	}

	// Create SaleFactorTax instance
	saleFactorTax := models.SaleFactorTax{
		SaleFactorConfirmationID: requestBody.SaleFactorConfirmationID,
		BillType:                 requestBody.BillType,
		PostType:                 requestBody.PostType,
		CreationDate:             requestBody.CreationDate,
		SettlementMethod:         requestBody.SettlementMethod,
		CashAmount:               requestBody.CashAmount,
		LoanAmount:               requestBody.LoanAmount,
	}

	// Save to database
	if err := db.Create(&saleFactorTax).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to create SaleFactorTax",
		})
	}

	return c.JSON(http.StatusCreated, saleFactorTax)
}

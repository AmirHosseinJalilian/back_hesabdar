package sale_factor_tax_status

import (
	"github.com/AmirHosseinJalilian/back_hesabdar/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func CreateSaleFactorTaxStatus(c echo.Context, db *gorm.DB) error {
	var requestBody models.CustomSaleFactorTaxStatus

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request body",
		})
	}

	// Create SaleFactorTaxStatus instance
	saleFactorTaxStatus := models.SaleFactorTaxStatus{
		SaleFactorConfirmationID: requestBody.SaleFactorConfirmationID,
		Status:                   requestBody.Status,
		StatusDate:               requestBody.StatusDate,
	}

	// Save to database
	if err := db.Create(&saleFactorTaxStatus).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to create SaleFactorTaxStatus",
		})
	}

	return c.JSON(http.StatusCreated, saleFactorTaxStatus)
}

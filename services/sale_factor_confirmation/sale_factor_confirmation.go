package sale_factor_confirmation

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AmirHosseinJalilian/back_hesabdar/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type QuerySaleFactorConfirmationsResponseType struct {
	StatusCode int `json:"statusCode"`
	Data       struct {
		Limit      int                             `json:"limit"`
		Offset     int                             `json:"offset"`
		Page       int                             `json:"page"`
		TotalRows  int                             `json:"totalRows"`
		TotalPages int                             `json:"totalPages"`
		Items      []models.SaleFactorConfirmation `json:"items"`
	} `json:"data"`
}

// GetSaleFactorConfirmations retrieves sale factor confirmations with pagination
func GetSaleFactorConfirmations(c echo.Context, db *gorm.DB) error {
	limitStr := c.QueryParam("limit")
	offsetStr := c.QueryParam("offset")
	pageStr := c.QueryParam("page")
	idStr := c.QueryParam("id")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10 // Default limit
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0 // Default offset
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1 // Default page
	}

	// Calculate offset based on page and limit
	if page > 1 {
		offset = (page - 1) * limit
	}

	var saleFactorConfirmations []models.SaleFactorConfirmation
	var totalRows int64

	query := db.Model(&models.SaleFactorConfirmation{}).
		Preload("Details.Commodity").                        // Preload Commodity in Details
		Preload("PepoleGrouping").                           // Preload PepoleGrouping
		Preload("PepoleGrouping.Pepoles").                   // Preload Pepoles in PepoleGrouping
		Preload("PepoleGrouping.Pepoles.PepoleDescriptions") // Preload PepoleDescriptions in Pepoles

	if idStr != "" {
		var id int64
		if id, err = strconv.ParseInt(idStr, 10, 64); err == nil {
			query = query.Where("id = ?", id)
		}
	}

	query.Count(&totalRows)
	query = query.Offset(offset).Limit(limit).Find(&saleFactorConfirmations)
	if query.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("Failed to execute query: %v", query.Error),
		})
	}

	// Calculate total pages
	totalPages := (int(totalRows) + limit - 1) / limit

	for i := range saleFactorConfirmations {
		saleFactorConfirmations[i].RowID = generateRowID(offset + i + 1)
	}

	responseData := QuerySaleFactorConfirmationsResponseType{
		StatusCode: http.StatusOK,
	}
	responseData.Data.Limit = limit
	responseData.Data.Offset = offset
	responseData.Data.Page = page
	responseData.Data.TotalRows = int(totalRows)
	responseData.Data.TotalPages = totalPages
	responseData.Data.Items = saleFactorConfirmations

	return c.JSON(http.StatusOK, responseData)
}

func generateRowID(index int) string {
	return fmt.Sprintf("%d", index)
}

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
		Limit      int                                   `json:"limit"`
		Offset     int                                   `json:"offset"`
		Page       int                                   `json:"page"`
		TotalRows  int                                   `json:"totalRows"`
		TotalPages int                                   `json:"totalPages"`
		Items      []models.CustomSaleFactorConfirmation `json:"items"`
	} `json:"data"`
}

// func generateRowID(index int) string {
// 	return fmt.Sprintf("%d", index)
// }

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
	// var SaleFactorConfirmationDetails []models.CustomSaleFactorConfirmationDetail
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

	// for i := range saleFactorConfirmations {
	// 	saleFactorConfirmations[i].RowID = generateRowID(offset + i + 1)
	// }

	// for i := range SaleFactorConfirmationDetails {
	// 	SaleFactorConfirmationDetails[i].DRowID = generateRowID(offset + i + 1)
	// }

	// Transform data into custom response struct
	var customSaleFactorConfirmations []models.CustomSaleFactorConfirmation
	for _, saleFactor := range saleFactorConfirmations {
		var customDetails []models.CustomSaleFactorConfirmationDetail
		for _, detail := range saleFactor.Details {
			customDetails = append(customDetails, models.CustomSaleFactorConfirmationDetail{
				ID: detail.ID,
				// DRowID:                   detail.DRowID,
				SaleFactorConfirmationID: detail.SaleFactorConfirmationID,
				Count:                    detail.Count,
				UnitCost:                 detail.UnitCost,
				CommodityDiscount:        detail.CommodityDiscount,
				ISCommodityDiscount:      detail.ISCommodityDiscount,
				Vat:                      detail.Vat,
				CommodityID:              detail.CommodityID,
				Commodity: models.CustomCommodity{
					ID:            detail.Commodity.ID,
					ComodityCod:   detail.Commodity.ComodityCod,
					CommodityName: detail.Commodity.CommodityName,
					UnitCount:     detail.Commodity.UnitCount,
					BasePrice:     detail.Commodity.BasePrice,
				},
			})
		}
		customSaleFactorConfirmations = append(customSaleFactorConfirmations, models.CustomSaleFactorConfirmation{
			ID: saleFactor.ID,
			// RowID:            saleFactor.RowID,
			DateFactorSale:   saleFactor.DateFactorSale,
			FactorNumber:     saleFactor.FactorNumber,
			SaleType:         saleFactor.SaleType,
			PepoleGroupingID: saleFactor.PepoleGroupingID,
			Details:          customDetails,
			PepoleGrouping: models.CustomPepoleGrouping{
				ID:          saleFactor.PepoleGrouping.ID,
				ObjectValue: saleFactor.PepoleGrouping.ObjectValue,
				Pepoles: func() []models.CustomPepole {
					var pepoles []models.CustomPepole
					for _, pepole := range saleFactor.PepoleGrouping.Pepoles {
						var descriptions []models.CustomPepoleDescription
						for _, desc := range pepole.PepoleDescriptions {
							descriptions = append(descriptions, models.CustomPepoleDescription{
								ID:              desc.ID,
								PepoleID:        desc.PepoleID,
								Address:         desc.Address,
								Phone:           desc.Phone,
								NationalityCode: desc.NationalityCode,
							})
						}
						pepoles = append(pepoles, models.CustomPepole{
							ID:                 pepole.ID,
							Name:               pepole.Name,
							PepoleType:         pepole.PepoleType,
							CodPepole:          pepole.CodPepole,
							GroupingID:         pepole.GroupingID,
							PepoleDescriptions: descriptions,
						})
					}
					return pepoles
				}(),
			},
		})
	}

	responseData := QuerySaleFactorConfirmationsResponseType{
		StatusCode: http.StatusOK,
	}
	responseData.Data.Limit = limit
	responseData.Data.Offset = offset
	responseData.Data.Page = page
	responseData.Data.TotalRows = int(totalRows)
	responseData.Data.TotalPages = totalPages
	responseData.Data.Items = customSaleFactorConfirmations

	return c.JSON(http.StatusOK, responseData)
}

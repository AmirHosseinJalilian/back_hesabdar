package sale_factor_confirmation_details

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb" // SQL Server driver
	"github.com/labstack/echo/v4"
)

// Define a struct for an invoice
type saleFactorConfirmationDetail struct {
	ID                       int64   `json:"id"`
	SaleFactorConfirmationID int64   `json:"saleFactorConfirmationID"`
	Commodity                int64   `json:"commodity"`
	Count                    float64 `json:"count"`
	UnitCost                 float64 `json:"unitCost"`
	CommodityDiscount        float64 `json:"commodityDiscount"`
	ISCommodityDiscount      bool    `json:"iSCommodityDiscount"`
	Vat                      float64 `json:"vat"`
	ComodityCod              string  `json:"comodityCod"`
	CommodityName            string  `json:"commodityName"`
	UnitCount                int64   `json:"unitCount"`
	BasePrice                int64   `json:"basePrice"`
}

// SQL Server connection parameters

func GetSaleFactorConfirmationDetails(c echo.Context, db *sql.DB) error {
	query := `SELECT sfcd.id, sfcd.saleFactorConfirmationID, sfcd.commodity, sfcd.count, sfcd.unitCost, sfcd.commodityDiscount,
	sfcd.iSCommodityDiscount, sfcd.vat, c.comodityCod, c.commodityName, c.unitCount, c.basePrice
	FROM SaleFactorConfirmationDetails sfcd
	INNER JOIN Commoditym c ON c.id = sfcd.commodity`
	rows, err := db.Query(query)
	if err != nil {
		// Log the error for debugging
		fmt.Println("Error executing query:", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("Failed to execute query: %v", err),
		})
	}
	defer rows.Close()

	var saleFactorConfirmationDetails []saleFactorConfirmationDetail
	for rows.Next() {
		var saleFactorCD saleFactorConfirmationDetail
		if err := rows.Scan(&saleFactorCD.ID, &saleFactorCD.SaleFactorConfirmationID, &saleFactorCD.Commodity, &saleFactorCD.Count,
			&saleFactorCD.UnitCost, &saleFactorCD.CommodityDiscount, &saleFactorCD.ISCommodityDiscount, &saleFactorCD.Vat,
			&saleFactorCD.ComodityCod, &saleFactorCD.CommodityName, &saleFactorCD.UnitCount, &saleFactorCD.BasePrice); err != nil {
			fmt.Println("Error scanning row:", err)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": fmt.Sprintf("Failed to scan row: %v", err),
			})
		}
		saleFactorConfirmationDetails = append(saleFactorConfirmationDetails, saleFactorCD)
	}

	if err := rows.Err(); err != nil {
		// Log the error for debugging
		fmt.Println("Row iteration error:", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("Row iteration error: %v", err),
		})
	}

	responseData := map[string]interface{}{
		"statusCode": http.StatusOK,
		"data":       saleFactorConfirmationDetails,
	}

	return c.JSON(http.StatusOK, responseData)
}

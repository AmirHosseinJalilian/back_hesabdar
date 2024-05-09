package sale_factor_confirmation

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// Define a struct for an invoice
type SaleFactorConfirmation struct {
	ID             int64     `json:"id"`
	DateFactorSale time.Time `json:"dateFactorSale"`
	FactorNumber   string    `json:"factorNumber"`
	SaleType       int       `json:"saleType"`
}

func GetSaleFactorConfirmations(c echo.Context, db *sql.DB) error {
	// db := database.Connect()
	//fasjklhd
	query := "SELECT id, dateFactorSale, factorNumber, saleType FROM SaleFactorConfirmation"
	rows, err := db.Query(query)
	if err != nil {
		// Log the error for debugging
		fmt.Println("Error executing query:", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("Failed to execute query: %v", err),
		})
	}
	defer rows.Close()

	var saleFactorConfirmations []SaleFactorConfirmation
	for rows.Next() {
		var saleFactorC SaleFactorConfirmation
		if err := rows.Scan(&saleFactorC.ID, &saleFactorC.DateFactorSale, &saleFactorC.FactorNumber, &saleFactorC.SaleType); err != nil {
			// Log the error for debugging
			fmt.Println("Error scanning row:", err)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": fmt.Sprintf("Failed to scan row: %v", err),
			})
		}
		saleFactorConfirmations = append(saleFactorConfirmations, saleFactorC)
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
		"data":       saleFactorConfirmations,
	}

	return c.JSON(http.StatusOK, responseData)
}

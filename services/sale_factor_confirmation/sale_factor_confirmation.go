package sale_factor_confirmation

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	// "github.com/AmirHosseinJalilian/back_hesabdar/custom/convert_date"
	"github.com/labstack/echo/v4"
)

// Define a struct for an invoice
type SaleFactorConfirmation struct {
	ID             int64     `json:"id"`
	RowID          string    `json:"rowId"` // Add row ID field
	DateFactorSale time.Time `json:"dateFactorSale"`
	FactorNumber   string    `json:"factorNumber"`
	SaleType       int       `json:"saleType"`
}

type QuerySaleFactorConfirmationsResponseType struct {
	StatusCode int `json:"statusCode"`
	Data       struct {
		Limit      int                      `json:"limit"`
		Offset     int                      `json:"offset"`
		Page       int                      `json:"page"`
		TotalRows  int                      `json:"totalRows"`
		TotalPages int                      `json:"totalPages"`
		Items      []SaleFactorConfirmation `json:"items"`
	} `json:"data"`
}

// Modify the GetSaleFactorConfirmations function to handle pagination
func GetSaleFactorConfirmations(c echo.Context, db *sql.DB) error {

	limitStr := c.QueryParam("limit")
	offsetStr := c.QueryParam("offset")
	pageStr := c.QueryParam("page")

	// Convert query parameters to integers
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

	// Fetch total rows
	var totalRows int
	err = db.QueryRow("SELECT COUNT(*) FROM SaleFactorConfirmation").Scan(&totalRows)
	if err != nil {
		fmt.Println("Error fetching total rows:", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("Failed to fetch total rows: %v", err),
		})
	}

	// Calculate total pages
	totalPages := (totalRows + limit - 1) / limit

	// Execute the query with limit and offset
	query := "SELECT id, dateFactorSale, factorNumber, saleType FROM SaleFactorConfirmation ORDER BY id DESC OFFSET @offset ROWS FETCH NEXT @limit ROWS ONLY"
	rows, err := db.Query(query, sql.Named("limit", limit), sql.Named("offset", offset))
	if err != nil {
		fmt.Println("Error executing query:", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("Failed to execute query: %v", err),
		})
	}
	defer rows.Close()

	// Parse rows into struct
	var saleFactorConfirmations []SaleFactorConfirmation
	for i := 0; rows.Next(); i++ {
		var saleFactorC SaleFactorConfirmation
		if err := rows.Scan(&saleFactorC.ID, &saleFactorC.DateFactorSale, &saleFactorC.FactorNumber, &saleFactorC.SaleType); err != nil {
			fmt.Println("Error scanning row:", err)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": fmt.Sprintf("Failed to scan row: %v", err),
			})
		}
		// Generate RowID for each item
		saleFactorC.RowID = generateRowID(offset + i + 1)

		saleFactorConfirmations = append(saleFactorConfirmations, saleFactorC)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Row iteration error:", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("Row iteration error: %v", err),
		})
	}

	// Create the response data
	responseData := QuerySaleFactorConfirmationsResponseType{
		StatusCode: http.StatusOK,
	}
	responseData.Data.Limit = limit
	responseData.Data.Offset = offset
	responseData.Data.Page = page
	responseData.Data.TotalRows = totalRows
	responseData.Data.TotalPages = totalPages
	responseData.Data.Items = saleFactorConfirmations

	return c.JSON(http.StatusOK, responseData)
}

func generateRowID(index int) string {
	// Concatenate the page number with the index to create a unique row ID
	return fmt.Sprintf("%d", index)
}

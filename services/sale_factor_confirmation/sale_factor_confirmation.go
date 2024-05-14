package sale_factor_confirmation

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	// "github.com/AmirHosseinJalilian/back_hesabdar/custom/convert_date"
	// "github.com/hoitek-go/govalidity"
	"github.com/labstack/echo/v4"
)

// Define a struct for an invoice
type SaleFactorConfirmation struct {
	ID                       int64     `json:"id"`
	RowID                    string    `json:"rowId"` // Add row ID field
	DateFactorSale           time.Time `json:"dateFactorSale"`
	FactorNumber             string    `json:"factorNumber"`
	SaleType                 int       `json:"saleType"`
	PepoleGroupingID         int64     `json:"pepoleGroupingId"` // Add PepoleGroupingID field
	ObjectValue              string    `json:"objectValue"`
	Name                     string    `json:"name"`
	NationalityCode          string    `json:"nationalityCode"`
	SaleFactorConfirmationID int64     `json:"saleFactorConfirmationID"`
	Commodity                int64     `json:"commodity"`
	Count                    float64   `json:"count"`
	UnitCost                 float64   `json:"unitCost"`
	CommodityDiscount        float64   `json:"commodityDiscount"`
	ISCommodityDiscount      bool      `json:"iSCommodityDiscount"`
	Vat                      float64   `json:"vat"`
	ComodityCod              string    `json:"comodityCod"`
	CommodityName            string    `json:"commodityName"`
	UnitCount                int64     `json:"unitCount"`
	BasePrice                int64     `json:"basePrice"`
	Phone                    string    `json:"phone"`
	Address                  string    `json:"address"`
	PepoleType               int16     `json:"pepoleType"`
}

// type ProductQueryRequestParams struct {
// 	ID      string                 `json:"id,omitempty"`
// 	Order   string                 `json:"order,omitempty"`
// 	OrderBy string                 `json:"order_by,omitempty"`
// 	Query   string                 `json:"query,omitempty"`
// 	Filters SaleFactorConfirmation `json:"filters,omitempty"`
// }

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
	query := `SELECT sfc.id, sfc.dateFactorSale, sfc.factorNumber, sfc.saleType,sfc.PepoleGroupingID, g.ObjectValue, p.name,pd.nationalityCode,
	sfcd.saleFactorConfirmationID, sfcd.commodity, sfcd.count, sfcd.unitCost, sfcd.commodityDiscount,sfcd.iSCommodityDiscount ,
	sfcd.vat, c.comodityCod, c.commodityName, c.unitCount, c.basePrice, pd.phone, pd.address, p.pepoleType
	FROM SaleFactorConfirmation sfc
	INNER JOIN SaleFactorConfirmationDetails sfcd ON sfc.ID = sfcd.SaleFactorConfirmationID
	INNER JOIN Commoditym c ON c.id = sfcd.commodity
	INNER JOIN Grouping g ON sfc.PepoleGroupingID = g.ID
	INNER JOIN Pepole p ON g.ID = p.ID
	INNER JOIN PepoleDescription pd ON p.ID = pd.PepoleID
	ORDER BY sfc.id DESC
	OFFSET @offset ROWS FETCH NEXT @limit ROWS ONLY
	WHERE sfc.id = @id`
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
		if err := rows.Scan(&saleFactorC.ID, &saleFactorC.DateFactorSale, &saleFactorC.FactorNumber, &saleFactorC.SaleType,
			&saleFactorC.PepoleGroupingID, &saleFactorC.ObjectValue, &saleFactorC.Name, &saleFactorC.NationalityCode,
			&saleFactorC.SaleFactorConfirmationID, &saleFactorC.Commodity, &saleFactorC.Count, &saleFactorC.UnitCost,
			&saleFactorC.CommodityDiscount, &saleFactorC.ISCommodityDiscount, &saleFactorC.Vat, &saleFactorC.ComodityCod,
			&saleFactorC.CommodityName, &saleFactorC.UnitCount, &saleFactorC.BasePrice, &saleFactorC.Phone,
			&saleFactorC.Address, &saleFactorC.PepoleType); err != nil {
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

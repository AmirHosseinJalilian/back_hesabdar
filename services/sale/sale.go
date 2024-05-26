package sale

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Commodity struct {
	ID            int64  `json:"id"`
	ComodityCod   string `json:"comodityCod"`
	CommodityName string `json:"commodityName"`
	UnitCount     int64  `json:"unitCount"`
	BasePrice     int64  `json:"basePrice"`
}

type SaleFactorDetail struct {
	ID                       int64     `json:"id"`
	SaleFactorConfirmationID int64     `json:"saleFactorConfirmationID"`
	Count                    float64   `json:"count"`
	UnitCost                 float64   `json:"unitCost"`
	CommodityDiscount        float64   `json:"commodityDiscount"`
	ISCommodityDiscount      bool      `json:"iSCommodityDiscount"`
	Vat                      float64   `json:"vat"`
	CommodityID              int64     `json:"commodityID"`
	Commodity                Commodity `json:"commodity"` // Single Commodity instead of []Commodity
}

type PepoleDescription struct {
	ID              int64  `json:"id"`
	PepoleID        int64  `json:"pepoleID"`
	Address         string `json:"address"`
	Phone           string `json:"phone"`
	NationalityCode string `json:"nationalityCode"`
}

type Pepole struct {
	ID                 int64               `json:"id"`
	Name               string              `json:"name"`
	PepoleType         int16               `json:"pepoleType"`
	CodPepole          string              `json:"codPepole"`
	GroupingID         int64               `json:"groupingID"`
	PepoleDescriptions []PepoleDescription `json:"pepoleDescriptions"` // Slice of PepoleDescription instead of single PepoleDescription
}

type Grouping struct {
	ID          int64    `json:"id"`
	ObjectValue string   `json:"objectValue"`
	Pepoles     []Pepole `json:"pepoles"` // Slice of Pepole instead of single Pepole
}

type SaleFactorConfirmation struct {
	ID               int64              `json:"id"`
	RowID            string             `json:"rowId"`
	DateFactorSale   time.Time          `json:"dateFactorSale"`
	FactorNumber     string             `json:"factorNumber"`
	SaleType         int                `json:"saleType"`
	PepoleGroupingID int64              `json:"pepoleGroupingID"`
	Details          []SaleFactorDetail `json:"Details"`
	PepoleGrouping   Grouping           `json:"PepoleGrouping"`
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

func generateRowID(rowNumber int) string {
	return fmt.Sprintf("row_%d", rowNumber)
}

func GetSale(c echo.Context, db *sql.DB) error {
	limitStr := c.QueryParam("limit")
	offsetStr := c.QueryParam("offset")
	pageStr := c.QueryParam("page")
	idStr := c.QueryParam("id")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	if page > 1 {
		offset = (page - 1) * limit
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		id = -1
	}

	var totalRows int
	queryTotal := "SELECT COUNT(*) FROM SaleFactorConfirmation"
	if id != -1 {
		queryTotal += " WHERE id = @id"
	}
	err = db.QueryRow(queryTotal, sql.Named("id", id)).Scan(&totalRows)
	if err != nil {
		fmt.Println("Error fetching total rows:", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("Failed to fetch total rows: %v", err),
		})
	}

	totalPages := (totalRows + limit - 1) / limit

	query := `SELECT 
        sfc.id, sfc.dateFactorSale, sfc.factorNumber, sfc.saleType, sfc.PepoleGroupingID,
        sfcd.id AS sfcd_id, sfcd.saleFactorConfirmationID, sfcd.count, sfcd.unitCost, sfcd.commodityDiscount,
        sfcd.iSCommodityDiscount, sfcd.vat, c.id AS commodityID, c.comodityCod, c.commodityName, c.unitCount, c.basePrice,
        p.id AS pepoleID, p.name AS pepoleName, p.pepoleType, pd.nationalityCode, pd.phone, pd.address,
        g.id AS groupingID, g.objectValue
    FROM SaleFactorConfirmation sfc
    INNER JOIN SaleFactorConfirmationDetails sfcd ON sfc.ID = sfcd.SaleFactorConfirmationID
    INNER JOIN Commoditym c ON c.id = sfcd.commodity
    INNER JOIN Grouping g ON sfc.PepoleGroupingID = g.ID
    INNER JOIN Pepole p ON g.ID = p.ID
    INNER JOIN PepoleDescription pd ON p.ID = pd.PepoleID
    WHERE (@id = -1 OR sfc.id = @id)
    ORDER BY sfc.id DESC
    OFFSET @offset ROWS
    FETCH NEXT @limit ROWS ONLY`

	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println("Error preparing query:", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("Failed to prepare query: %v", err),
		})
	}
	defer stmt.Close()

	rows, err := stmt.Query(sql.Named("limit", limit), sql.Named("offset", offset), sql.Named("id", id))
	if err != nil {
		fmt.Println("Error executing query:", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("Failed to execute query: %v", err),
		})
	}
	defer rows.Close()

	saleFactorConfirmationsMap := make(map[int64]*SaleFactorConfirmation)
	saleFactorDetailsMap := make(map[int64][]SaleFactorDetail)
	pepoleDescriptionsMap := make(map[int64][]PepoleDescription)

	for rows.Next() {
		var sfc SaleFactorConfirmation
		var sfcd SaleFactorDetail
		var commodity Commodity
		var pepole Pepole
		var pepoleDescription PepoleDescription
		var grouping Grouping

		if err := rows.Scan(
			&sfc.ID, &sfc.DateFactorSale, &sfc.FactorNumber, &sfc.SaleType, &sfc.PepoleGroupingID,
			&sfcd.ID, &sfcd.SaleFactorConfirmationID, &sfcd.Count, &sfcd.UnitCost, &sfcd.CommodityDiscount,
			&sfcd.ISCommodityDiscount, &sfcd.Vat, &commodity.ID, &commodity.ComodityCod, &commodity.CommodityName,
			&commodity.UnitCount, &commodity.BasePrice, &pepole.ID, &pepole.Name, &pepole.PepoleType,
			&pepoleDescription.NationalityCode, &pepoleDescription.Phone, &pepoleDescription.Address,
			&grouping.ID, &grouping.ObjectValue,
		); err != nil {
			fmt.Println("Error scanning row:", err)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": fmt.Sprintf("Failed to scan row: %v", err),
			})
		}

		sfcd.CommodityID = commodity.ID
		sfcd.Commodity = commodity

		pepoleDescription.PepoleID = pepole.ID
		pepoleDescriptionsMap[pepole.ID] = append(pepoleDescriptionsMap[pepole.ID], pepoleDescription)

		if _, exists := saleFactorConfirmationsMap[sfc.ID]; !exists {
			sfc.RowID = generateRowID(offset + len(saleFactorConfirmationsMap) + 1)
			saleFactorConfirmationsMap[sfc.ID] = &sfc
		}

		saleFactorDetailsMap[sfc.ID] = append(saleFactorDetailsMap[sfc.ID], sfcd)
	}

	var saleFactorConfirmations []SaleFactorConfirmation
	for _, sfc := range saleFactorConfirmationsMap {
		sfc.Details = saleFactorDetailsMap[sfc.ID]
		var pepoles []Pepole
		for pepoleID, pepoleDescriptions := range pepoleDescriptionsMap {
			var pepole Pepole
			pepole.ID = pepoleID
			pepole.PepoleDescriptions = pepoleDescriptions
			pepoles = append(pepoles, pepole)
		}
		sfc.PepoleGrouping.Pepoles = pepoles
		saleFactorConfirmations = append(saleFactorConfirmations, *sfc)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Row iteration error:", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("Row iteration error: %v", err),
		})
	}

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

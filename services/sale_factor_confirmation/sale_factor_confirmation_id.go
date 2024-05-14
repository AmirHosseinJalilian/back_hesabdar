package sale_factor_confirmation

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	// "github.com/AmirHosseinJalilian/back_hesabdar/custom/convert_date"
	// "github.com/hoitek-go/govalidity"
	"github.com/labstack/echo/v4"
)

func GetSaleFactorConfirmationByID(c echo.Context, db *sql.DB) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid SaleFactorConfirmation ID",
		})
	}

	var saleFactorC SaleFactorConfirmation
	query := `
        SELECT sfc.id, sfc.dateFactorSale, sfc.factorNumber, sfc.saleType, sfc.PepoleGroupingID, g.ObjectValue, p.name, pd.nationalityCode,
               sfcd.saleFactorConfirmationID, sfcd.commodity, sfcd.count, sfcd.unitCost, sfcd.commodityDiscount, sfcd.iSCommodityDiscount,
               sfcd.vat, c.comodityCod, c.commodityName, c.unitCount, c.basePrice, pd.phone, pd.address, p.pepoleType
        FROM SaleFactorConfirmation sfc
        INNER JOIN SaleFactorConfirmationDetails sfcd ON sfc.ID = sfcd.SaleFactorConfirmationID
        INNER JOIN Commoditym c ON c.id = sfcd.commodity
        INNER JOIN Grouping g ON sfc.PepoleGroupingID = g.ID
        INNER JOIN Pepole p ON g.ID = p.ID
        INNER JOIN PepoleDescription pd ON p.ID = pd.PepoleID
        WHERE sfc.id = @id
    `
	row := db.QueryRow(query, sql.Named("id", id))
	err = row.Scan(&saleFactorC.ID, &saleFactorC.DateFactorSale, &saleFactorC.FactorNumber, &saleFactorC.SaleType,
		&saleFactorC.PepoleGroupingID, &saleFactorC.ObjectValue, &saleFactorC.Name, &saleFactorC.NationalityCode,
		&saleFactorC.SaleFactorConfirmationID, &saleFactorC.Commodity, &saleFactorC.Count, &saleFactorC.UnitCost,
		&saleFactorC.CommodityDiscount, &saleFactorC.ISCommodityDiscount, &saleFactorC.Vat, &saleFactorC.ComodityCod,
		&saleFactorC.CommodityName, &saleFactorC.UnitCount, &saleFactorC.BasePrice, &saleFactorC.Phone,
		&saleFactorC.Address, &saleFactorC.PepoleType)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"error": "SaleFactorConfirmation not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("Failed to fetch SaleFactorConfirmation: %v", err),
		})
	}

	return c.JSON(http.StatusOK, saleFactorC)
}

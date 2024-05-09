package pepole

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Define a struct for an invoice
type pepole struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	CodPepole       string `json:"codPepole"`
	PepoleType      uint8  `json:"pepoleType"`
	PepoleID        int    `json:"pepoleID"`
	Address         string `json:"address"`
	Phone           string `json:"phone"`
	NationalityCode string `json:"nationalityCode"`
}

func GetPepoles(c echo.Context, db *sql.DB) error {
	query := `SELECT p.id, p.name, p.codPepole, p.pepoleType, pd.pepoleID, pd.address, pd.phone, pd.nationalityCode FROM Pepole p
	INNER JOIN PepoleDescription pd ON p.id = pd.pepoleID`
	rows, err := db.Query(query)
	if err != nil {
		// Log the error for debugging
		fmt.Println("Error executing query:", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("Failed to execute query: %v", err),
		})
	}
	defer rows.Close()

	var Pepoles []pepole
	for rows.Next() {
		var pepole pepole
		if err := rows.Scan(&pepole.ID, &pepole.Name, &pepole.CodPepole, &pepole.PepoleType, &pepole.PepoleID, &pepole.Address,
			&pepole.Phone, &pepole.NationalityCode); err != nil {
			// Log the error for debugging
			fmt.Println("Error scanning row:", err)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": fmt.Sprintf("Failed to scan row: %v", err),
			})
		}
		Pepoles = append(Pepoles, pepole)
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
		"data":       Pepoles,
	}
	return c.JSON(http.StatusOK, responseData)

}

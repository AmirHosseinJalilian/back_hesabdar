package main

import (
	"fmt"
	"net/http"

	"github.com/AmirHosseinJalilian/back_hesabdar/database"
	"github.com/AmirHosseinJalilian/back_hesabdar/services/sale_factor_confirmation"
	"github.com/labstack/echo/v4"
	"github.com/rs/cors"
)

func main() {
	e := echo.New()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allow all origins
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		Debug:          true,
	})
	e.Use(echo.WrapMiddleware(c.Handler))

	e.GET("/SaleFactorConfirmations", func(c echo.Context) error {
		// host := c.QueryParam("host")
		// port := c.QueryParam("port")
		user := c.QueryParam("user")
		password := c.QueryParam("password")
		dbName := c.QueryParam("dbName")

		if user == "" || password == "" || dbName == "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "Missing required query parameters",
			})
		}

		db, err := database.Connect(user, password, dbName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": fmt.Sprintf("Failed to connect to database: %v", err),
			})
		}

		return sale_factor_confirmation.GetSaleFactorConfirmations(c, db)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

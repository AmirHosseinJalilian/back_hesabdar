package main

import (
	"github.com/AmirHosseinJalilian/back_hesabdar/database"
	"github.com/AmirHosseinJalilian/back_hesabdar/services/pepole"
	"github.com/AmirHosseinJalilian/back_hesabdar/services/sale_factor_confirmation"
	"github.com/AmirHosseinJalilian/back_hesabdar/services/sale_factor_confirmation_details"
	"github.com/labstack/echo/v4"
	"github.com/rs/cors"
)

func main() {
	// Echo instance
	e := echo.New()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allow all origins
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		Debug:          true,
	})
	// Wrap Echo instance with CORS handler
	e.Use(echo.WrapMiddleware(c.Handler))
	// Connect to the database
	db := database.Connect()
	// Remember to defer closing the database connection until application stops
	defer db.Close()
	// Define routes
	e.GET("/Pepoles", func(c echo.Context) error {
		return pepole.GetPepoles(c, db)
	})
	e.GET("/SaleFactorConfirmations", func(c echo.Context) error {
		return sale_factor_confirmation.GetSaleFactorConfirmations(c, db)
	})
	e.GET("/SaleFactorConfirmationDetails", func(c echo.Context) error {
		return sale_factor_confirmation_details.GetSaleFactorConfirmationDetails(c, db)
	})
	e.Logger.Fatal(e.Start(":8080"))
}

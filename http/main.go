package main

import (
	"github.com/AmirHosseinJalilian/back_hesabdar/database"
	"github.com/AmirHosseinJalilian/back_hesabdar/models"
	"github.com/AmirHosseinJalilian/back_hesabdar/services/sale_factor_confirmation"
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

	// Connect to the database using GORM
	db := database.Connect()
	sqlDB, err := db.DB()
	if err != nil {
		e.Logger.Fatal("Failed to get generic database object:", err)
	}
	defer sqlDB.Close()

	// Auto migrate the models
	db.AutoMigrate(&models.SaleFactorConfirmation{}, &models.SaleFactorConfirmationDetails{}, &models.Commoditym{}, &models.Grouping{}, &models.Pepole{}, &models.PepoleDescription{})

	// Define routes
	e.GET("/SaleFactorConfirmations", func(c echo.Context) error {
		return sale_factor_confirmation.GetSaleFactorConfirmations(c, db)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

// import (
// 	"github.com/AmirHosseinJalilian/back_hesabdar/database"
// 	"github.com/AmirHosseinJalilian/back_hesabdar/services/sale"
// 	"github.com/labstack/echo/v4"
// 	"github.com/rs/cors"
// )

// func main() {
// 	// Echo instance
// 	e := echo.New()
// 	c := cors.New(cors.Options{
// 		AllowedOrigins: []string{"*"}, // Allow all origins
// 		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
// 		AllowedHeaders: []string{"*"},
// 		Debug:          true,
// 	})
// 	// Wrap Echo instance with CORS handler
// 	e.Use(echo.WrapMiddleware(c.Handler))
// 	// Connect to the database
// 	db := database.Connect()
// 	// Remember to defer closing the database connection until application stops
// 	defer db.Close()
// 	// Define routes

// 	e.GET("/Sale", func(c echo.Context) error {
// 		return sale.GetSale(c, db)
// 	})

// 	e.Logger.Fatal(e.Start(":8080"))
// }

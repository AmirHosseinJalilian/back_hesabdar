package main

import (
	// "log"
	// "github.com/AmirHosseinJalilian/back_hesabdar/services/pepole"
	// "github.com/gin-gonic/gin"
	"github.com/AmirHosseinJalilian/back_hesabdar/database"
	"github.com/AmirHosseinJalilian/back_hesabdar/services/sale_factor_confirmation"
	"github.com/AmirHosseinJalilian/back_hesabdar/services/sale_factor_confirmation_details"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "github.com/rs/cors"
	// "github.com/AmirHosseinJalilian/back_hesabdar/services/pepole"
	// "github.com/AmirHosseinJalilian/back_hesabdar/services/sale_factor_confirmation"
	// "github.com/AmirHosseinJalilian/back_hesabdar/services/sale_factor_confirmation_details"
)

func main() {
	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Connect to the database
	db := database.Connect()
	// Remember to defer closing the database connection until application stops
	defer db.Close()
	// Routes
	// Define routes
	e.GET("/SaleFactorConfirmations", func(c echo.Context) error {
		return sale_factor_confirmation.GetSaleFactorConfirmations(c, db)
	})
	e.GET("/SaleFactorConfirmationDetails", func(c echo.Context) error {
		return sale_factor_confirmation_details.GetSaleFactorConfirmationDetails(c, db)
	})
	e.Start(":8080")
	// log.Fatal(http.ListenAndServe(serverPort, router))

	// 	mux := http.NewServeMux()
	// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 		w.Header().Set("Content-Type", "application/json")
	// 		w.Write([]byte("{\"hello\": \"world\"}"))
	// 	})

	// 	// cors.Default() setup the middleware with default options being
	// 	// all origins accepted with simple methods (GET, POST). See
	// 	// documentation below for more options.
	// 	handler := cors.Default().Handler(mux)
	// 	http.ListenAndServe(":8080", handler)
}

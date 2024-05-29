package login

import (
	"net/http"

	"github.com/AmirHosseinJalilian/back_hesabdar/database"
	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	DBName   string `json:"dbname" form:"dbname"`
}

func Login(c echo.Context) error {
	loginRequest := new(LoginRequest)
	if err := c.Bind(loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	// Use the provided credentials and dbName to attempt a connection
	db, err := database.Connect("192.168.1.109", "7007", loginRequest.Username, loginRequest.Password, loginRequest.DBName)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid credentials or database name",
		})
	}

	// Successfully connected to the database
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully connected to the database",
	})
}

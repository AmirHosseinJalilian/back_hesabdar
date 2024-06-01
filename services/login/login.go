package login

import (
	"net/http"

	"github.com/AmirHosseinJalilian/back_hesabdar/database"
	"github.com/AmirHosseinJalilian/back_hesabdar/models"
	"github.com/labstack/echo/v4"
	// "gorm.io/gorm"
)

func Login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")
	dbName := c.QueryParam("dbname")

	if username == "" || password == "" || dbName == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Missing required query parameters",
		})
	}

	db, err := database.Connect(username, password, dbName)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid credentials or database name",
		})
	}

	// Close the database connection
	sqlDB, err := db.DB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to get the underlying *sql.DB",
		})
	}
	defer sqlDB.Close()

	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid username or password",
		})
	}

	if user.Password != password {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid username or password",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully connected to the database",
	})
}

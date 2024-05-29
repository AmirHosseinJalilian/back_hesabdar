package database

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func Connect(user, password, dbName string) (*gorm.DB, error) {

	const (
		host = "192.168.1.109"
		port = "7007"
	)

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		user, password, host, port, dbName)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open connection to SQL Server database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get generic database object: %w", err)
	}
	if err = sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping SQL Server database: %w", err)
	}

	fmt.Println("Connected to SQL Server database using GORM.")
	return db, nil
}

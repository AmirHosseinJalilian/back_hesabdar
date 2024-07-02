package database

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	const (
		host     = "192.168.1.109"
		port     = "7007"
		user     = "netim"     // Replace with your SQL Server username
		password = "smj920123" // Replace with your SQL Server password
		dbName   = "Mehrad"    // Replace with your SQL Server database name
	)

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		user, password, host, port, dbName)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open connection to SQL Server database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get generic database object: %v", err)
	}
	if err = sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping SQL Server database: %v", err)
	}

	fmt.Println("Connected to SQL Server database using GORM.")
	return db
}

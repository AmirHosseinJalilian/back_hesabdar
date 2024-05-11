package database

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

func Connect() *sql.DB {

	const (
		driverName = "sqlserver"
		host       = "192.168.1.18"
		// host     = "192.168.1.100"
		port     = "7007"
		user     = "netim"     // Replace with your SQL Server username
		password = "smj920123" // Replace with your SQL Server password
		dbName   = "Mehrad"    // Replace with your SQL Server database name
	)

	var db *sql.DB

	// Connect to the SQL Server
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		user, password, host, port, dbName)
	var err error
	db, err = sql.Open(driverName, dsn)
	if err != nil {
		log.Fatalf("Failed to open connection to SQL Server database: %v", err)
	}
	// defer db.Close()

	// Verify connection
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping SQL Server database: %v", err)
	}
	fmt.Println("Connected to SQL Server database.")

	// Create a new router
	// router := mux.NewRouter()

	// Define routes
	// router.HandleFunc("/SaleFactorConfirmations", getSaleFactorConfirmations).Methods("GET", "OPTIONS")

	// Start the server
	serverPort := ":8080"
	fmt.Printf("Server started on port %s\n", serverPort)
	if err != nil {
		log.Fatal(err)
	}
	// log.Fatal(http.ListenAndServe(serverPort, router))

	return db
}

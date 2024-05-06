package sale_factor_confirmation

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/denisenkom/go-mssqldb" // SQL Server driver
	"github.com/gorilla/mux"
)

// Define a struct for an invoice
type saleFactorConfirmation struct {
	ID             int64     `json:"id"`
	DateFactorSale time.Time `json:"dateFactorSale"`
	FactorNumber   string    `json:"factorNumber"`
	SaleType       int       `json:"saleType"`
	// PepoleGroupingID int    `json:"pepoleID"`
	// Address         string `json:"address"`
	// Phone           string `json:"phone"`
	// NationalityCode string `json:"nationalityCode"`
}

// SQL Server connection parameters
const (
	driverName = "sqlserver"
	host       = "192.168.1.18"
	port       = "7007"
	user       = "netim"     // Replace with your SQL Server username
	password   = "smj920123" // Replace with your SQL Server password
	dbName     = "Mehrad"    // Replace with your SQL Server database name
)

var db *sql.DB

func SaleFactorConfirmation() {
	// Connect to the SQL Server
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		user, password, host, port, dbName)
	var err error
	db, err = sql.Open(driverName, dsn)
	if err != nil {
		log.Fatalf("Failed to open connection to SQL Server database: %v", err)
	}
	defer db.Close()

	// Verify connection
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping SQL Server database: %v", err)
	}
	fmt.Println("Connected to SQL Server database.")

	// Create a new router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/SaleFactorConfirmation", getSaleFactorConfirmations).Methods("GET")

	// Start the server
	// serverAddress := "your-site-address:your-port"
	serverPort := ":8080"
	fmt.Printf("Server started on port %s\n", serverPort)
	log.Fatal(http.ListenAndServe(serverPort, router))
}

func getSaleFactorConfirmations(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, dateFactorSale, factorNumber, saleType FROM SaleFactorConfirmation"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to execute query: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var saleFactorConfirmations []saleFactorConfirmation
	for rows.Next() {
		var saleFactorC saleFactorConfirmation
		if err := rows.Scan(&saleFactorC.ID, &saleFactorC.DateFactorSale, &saleFactorC.FactorNumber, &saleFactorC.SaleType); err != nil {
			http.Error(w, fmt.Sprintf("Failed to scan row: %v", err), http.StatusInternalServerError)
			return
		}
		saleFactorConfirmations = append(saleFactorConfirmations, saleFactorC)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Row iteration error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(saleFactorConfirmations)
}

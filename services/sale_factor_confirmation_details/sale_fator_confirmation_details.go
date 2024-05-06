package sale_factor_confirmation_details

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb" // SQL Server driver
	"github.com/gorilla/mux"
)

// Define a struct for an invoice
type saleFactorConfirmationDetail struct {
	ID                       int64   `json:"id"`
	SaleFactorConfirmationID int64   `json:"saleFactorConfirmationID"`
	Commodity                int64   `json:"commodity"`
	Count                    float64 `json:"count"`
	UnitCost                 float64 `json:"unitCost"`
	CommodityDiscount        float64 `json:"commodityDiscount"`
	ISCommodityDiscount      bool    `json:"iSCommodityDiscount"`
	Vat                      float64 `json:"vat"`
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

func SaleFactorConfirmationDetails() {
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
	router.HandleFunc("/SaleFactorConfirmationDetails", getSaleFactorConfirmationDetails).Methods("GET")

	// Start the server
	// serverAddress := "your-site-address:your-port"
	serverPort := ":8080"
	fmt.Printf("Server started on port %s\n", serverPort)
	log.Fatal(http.ListenAndServe(serverPort, router))
}

func getSaleFactorConfirmationDetails(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, saleFactorConfirmationID, commodity, count, unitCost, commodityDiscount, iSCommodityDiscount, vat FROM SaleFactorConfirmationDetails"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to execute query: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var saleFactorConfirmationDetails []saleFactorConfirmationDetail
	for rows.Next() {
		var saleFactorCD saleFactorConfirmationDetail
		if err := rows.Scan(&saleFactorCD.ID, &saleFactorCD.SaleFactorConfirmationID, &saleFactorCD.Commodity, &saleFactorCD.Count,
			&saleFactorCD.UnitCost, &saleFactorCD.CommodityDiscount, &saleFactorCD.ISCommodityDiscount, &saleFactorCD.Vat); err != nil {
			http.Error(w, fmt.Sprintf("Failed to scan row: %v", err), http.StatusInternalServerError)
			return
		}
		saleFactorConfirmationDetails = append(saleFactorConfirmationDetails, saleFactorCD)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Row iteration error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(saleFactorConfirmationDetails)
}

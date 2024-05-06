package pepole

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
type pepole struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	CodPepole       string `json:"codPepole"`
	PepoleType      uint8  `json:"pepoleType"`
	PepoleID        int    `json:"pepoleID"`
	Address         string `json:"address"`
	Phone           string `json:"phone"`
	NationalityCode string `json:"nationalityCode"`
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

func Pepole() {
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
	router.HandleFunc("/Pepoles", getPepoles).Methods("GET")

	// Start the server
	// serverAddress := "your-site-address:your-port"
	serverPort := ":8080"
	fmt.Printf("Server started on port %s\n", serverPort)
	log.Fatal(http.ListenAndServe(serverPort, router))
}

func getPepoles(w http.ResponseWriter, r *http.Request) {
	query := `SELECT p.id, p.name, p.codPepole, p.pepoleType, pd.pepoleID, pd.address, pd.phone, pd.nationalityCode FROM Pepole p
	INNER JOIN PepoleDescription pd ON p.id = pd.pepoleID`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to execute query: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var Pepoles []pepole
	for rows.Next() {
		var pepole pepole
		if err := rows.Scan(&pepole.ID, &pepole.Name, &pepole.CodPepole, &pepole.PepoleType, &pepole.PepoleID, &pepole.Address,
			&pepole.Phone, &pepole.NationalityCode); err != nil {
			http.Error(w, fmt.Sprintf("Failed to scan row: %v", err), http.StatusInternalServerError)
			return
		}
		Pepoles = append(Pepoles, pepole)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Row iteration error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Pepoles)
}

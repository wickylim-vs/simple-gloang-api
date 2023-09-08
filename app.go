package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Read MySQL connection parameters from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Construct the MySQL connection string
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Initialize the MySQL database connection
	var err error
	db, err = sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Set some database connection pool settings
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Create the "users" table if it doesn't exist
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL
		);
	`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	// Insert initial data into the "users" table if it's empty
	countRowsSQL := "SELECT COUNT(*) FROM users"
	var count int
	err = db.QueryRow(countRowsSQL).Scan(&count)
	if err != nil {
		log.Fatalf("Failed to count rows: %v", err)
	}

	if count == 0 {
		insertDataSQL := `
			INSERT INTO users (name, email) VALUES
			('John Doe', 'john@example.com'),
			('Jane Smith', 'jane@example.com');
		`
		_, err = db.Exec(insertDataSQL)
		if err != nil {
			log.Fatalf("Failed to insert initial data: %v", err)
		}
	}

	// Create a new router instance
	r := mux.NewRouter()

	// Define your API routes
	r.HandleFunc("/api/users", getUsers).Methods("GET")

	// Start the HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	serverAddr := fmt.Sprintf("0.0.0.0:%s", port)
	fmt.Printf("Listening on %s...\n", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, r))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	// Perform a database query to retrieve users
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			http.Error(w, "Failed to scan user data", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	// Convert the user data to JSON and send it as the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}


// cmd/setup_database.go
package cmd

import (
	"InvoiceMaster/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var setupDatabaseCmd = &cobra.Command{
	Use:   "setup_database",
	Short: "Setup the database and create the migrations table",
	Run: func(cmd *cobra.Command, args []string) {
		setupDatabase()
	},
}

func init() {
	rootCmd.AddCommand(setupDatabaseCmd)
}

func setupDatabase() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	if err != nil {
		log.Fatalf("Failed to read database configuration: %v", err)
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS database_migrations (
		id SERIAL PRIMARY KEY,
		filename VARCHAR(255) UNIQUE NOT NULL,
		applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Fatalf("Failed to create database_migrations table: %v", err)
	}

	fmt.Println("Database setup complete and database_migrations table created.")
}

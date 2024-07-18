package cmd

import (
	"InvoiceMaster/config"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Apply migrations sequentially",
	Run: func(cmd *cobra.Command, args []string) {
		migrate()
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func migrate() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer db.Close()

	// Ensure migrations table exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS database_migrations (
        id SERIAL PRIMARY KEY,
        filename TEXT UNIQUE,
        applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`)
	if err != nil {
		log.Fatalf("Failed to create database_migrations table: %v", err)
	}

	// Get already applied migrations
	appliedMigrations, err := getAppliedMigrations(db)
	if err != nil {
		log.Fatalf("Failed to get applied migrations: %v", err)
	}

	// Read migration files
	migrations, err := findMigrations("db/migration")
	if err != nil {
		log.Fatalf("Failed to find migration files: %v", err)
	}

	// Sort migrations by filename (timestamp)
	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Timestamp < migrations[j].Timestamp
	})

	// Apply migrations sequentially
	for _, migration := range migrations {
		if _, applied := appliedMigrations[migration.Filename]; applied {
			fmt.Printf("Skipping already applied migration: %s\n", migration.Filename)
			continue
		}

		err := applyMigration(db, migration.Filename)
		if err != nil {
			log.Fatalf("Failed to apply migration %s: %v", migration.Filename, err)
		}

		// Record applied migration in database
		_, err = db.Exec(`INSERT INTO database_migrations (filename) VALUES ($1)`, migration.Filename)
		if err != nil {
			log.Fatalf("Failed to record applied migration %s: %v", migration.Filename, err)
		}
		fmt.Printf("Applied migration: %s\n", migration.Filename)
	}
}

type migrationInfo struct {
	Filename  string
	Timestamp int64
}

func getAppliedMigrations(db *sql.DB) (map[string]bool, error) {
	rows, err := db.Query(`SELECT filename FROM database_migrations`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	appliedMigrations := make(map[string]bool)
	for rows.Next() {
		var filename string
		if err := rows.Scan(&filename); err != nil {
			return nil, err
		}
		appliedMigrations[filename] = true
	}

	return appliedMigrations, nil
}

func findMigrations(dir string) ([]migrationInfo, error) {
	var migrations []migrationInfo

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".sql" {
			filename := file.Name()
			parts := strings.Split(filename, "_")
			if len(parts) < 2 {
				continue // Skip invalid filenames
			}

			timestampStr := parts[0]
			timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
			if err != nil {
				continue // Skip files with invalid timestamps
			}

			migrations = append(migrations, migrationInfo{
				Filename:  filename,
				Timestamp: timestamp,
			})
		}
	}

	return migrations, nil
}

func applyMigration(db *sql.DB, filename string) error {
	data, err := os.ReadFile(filepath.Join("db/migration", filename))
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(string(data))
	if err != nil {
		return err
	}

	return tx.Commit()
}

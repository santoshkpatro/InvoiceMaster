package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate_migration [message]",
	Short: "Generate a new SQL migration file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		generateMigration(args[0])
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func generateMigration(message string) {
	timestamp := time.Now().Unix()
	filename := fmt.Sprintf("db/migration/%d_%s.sql", timestamp, message)

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Failed to create migration file: %v\n", err)
		return
	}
	defer file.Close()

	fmt.Printf("Created migration file: %s\n", filename)
}

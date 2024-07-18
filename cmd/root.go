package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Invoice Master",
	Short: "Invoice Master CLI",
	Long:  "A Invoice Master CLI Application for managing database.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

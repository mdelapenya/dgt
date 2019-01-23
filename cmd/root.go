package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dgt",
	Short: "dgt allows you to gather information about Spanish car plates",
	Long: `A Fast and Flexible CLI for gathering Spain's car plates for Eco stickers,
				built with love by mdelapenya and friends in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

// Execute execute root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed!")
	}
}

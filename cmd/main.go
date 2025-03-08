package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "app",
		Short: "A multi-command application",
		Long:  `This application demonstrates how to use Cobra for multi-command applications.`,
	}

	// Define a 'api' command
	apiCmd := &cobra.Command{
		Use:   "api",
		Short: "Start the api",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Api started")
		},
	}

	// Define a 'consumer' command
	consumerCmd := &cobra.Command{
		Use:   "consumer",
		Short: "Start the consumer",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Consumer started")
		},
	}

	// Define a 'analytic' command
	analyticCmd := &cobra.Command{
		Use:   "analytic",
		Short: "Start the analytic",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Analytic started")
		},
	}

	// Define a 'migration' command
	migrationCmd := &cobra.Command{
		Use:   "migration",
		Short: "Start the migration",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Migration started")
		},
	}

	// Add subcommands to the root command
	rootCmd.AddCommand(apiCmd, consumerCmd, analyticCmd, migrationCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

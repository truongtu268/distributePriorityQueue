package main

import (
	"fmt"
	"log"
	"os"

	"github.com/truongtu268/distributePriorityQueue/consumer"
	"github.com/truongtu268/distributePriorityQueue/external"
	"github.com/truongtu268/distributePriorityQueue/server"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "app",
		Short: "A multi-command application",
		Long:  `This application demonstrates how to use Cobra for multi-command applications.`,
	}
	redisHost := "localhost:6379"
	pgConn := "postgresql://admin:admin@localhost:5432/advertisement?sslmode=disable"

	// Define a 'api' command
	apiCmd := &cobra.Command{
		Use:   "api",
		Short: "Start the api",
		Run: func(cmd *cobra.Command, args []string) {
			api, err := server.NewAdServer(pgConn, redis.Options{Addr: redisHost})
			if err != nil {
				log.Fatal(err)
			}
			api.Run()
		},
	}

	// Define a 'consumer' command
	consumerCmd := &cobra.Command{
		Use:   "consumer",
		Short: "Start the consumer",
		Run: func(cmd *cobra.Command, args []string) {
			consumer.Execute(pgConn, redisHost)
		},
	}

	// Define a 'analytic' command
	analyticCmd := &cobra.Command{
		Use:   "analytic",
		Short: "Start the analytic",
		Run: func(cmd *cobra.Command, args []string) {
			external.Run()
		},
	}

	// Add subcommands to the root command
	rootCmd.AddCommand(apiCmd, consumerCmd, analyticCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

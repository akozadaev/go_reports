package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	envPath := "."
	envFileName := ".env"

	fullPath := envPath + "/" + envFileName

	if err := godotenv.Overload(fullPath); err != nil {
		log.Printf("[ERROR] failed with %+v", "No .env file found")
	}
}

func main() {
	commandName := "run:root:cmd"
	rootCmd := &cobra.Command{
		Use: commandName,
	}

	subCmd := &cobra.Command{
		Use:   "use",
		Short: "Use sub cmd",
		Long:  "Use sub cmd",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Executed sub cmd")
		},
	}

	subCmdGenerate := &cobra.Command{
		Use:   "generate:report",
		Short: "Generate report",
		Long:  "Generate report",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Executed sub cmd")
		},
	}
	rootCmd.AddCommand(subCmd)
	rootCmd.AddCommand(subCmdGenerate)
	rootCmd.Execute()
}

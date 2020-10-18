package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var greetingStr string

func init()  {
	greeting.Flags().StringVarP(&greetingStr, "greeting", "g", "", "print greeting")
	rootCmd.AddCommand(greeting)
}

var greeting = &cobra.Command{
	Use: "greeting",
	Short: "Print greeting.",
	Long: "Pring Greeting.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("hello, %s\n", greetingStr)
	},
}

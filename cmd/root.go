package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "expense-tracker",
	Short: "simple expense tracker to manage your finances.",
	Long: `simple expense tracker application to manage your finances. The application should allow users to add, delete, and view their expenses.
	Respository: httsp://github.com/V1nic1us/expense-tracker-go`,
	Args: cobra.MinimumNArgs(1),
  }
  

  func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
  }
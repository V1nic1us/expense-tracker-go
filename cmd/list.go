package cmd

import (
	"fmt"
	"os"
	"github.com/xuri/excelize/v2"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all expenses",
	Long:  `list all expenses from the list of expenses`,

	Run: func(cmd *cobra.Command, args []string) {
		var f *excelize.File
		var err error
		if _, err := os.Stat(FileName); err == nil {
			f, err = excelize.OpenFile(FileName)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		rows, err := f.GetRows("Sheet1")
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, row := range rows {
			fmt.Printf("%-5s %-15s %-20s %-10s\n", row[0], row[1], row[2], row[3])
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
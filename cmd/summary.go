package cmd

import (
	"fmt"
	"os"
	"strings"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
)

var Month int

var sumaryCmd = &cobra.Command{
	Use:   "sumary",
	Short: "Users can view a summary of all expenses",
	Long:  "Users can view a summary of expenses for a specific month (of current year)",

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

		rows, err := f.GetRows(SheetName)
		if err != nil {
			fmt.Println(err)
			return
		}
		for i, row := range rows {
			if i == 0 {
				fmt.Printf("%-5s %-15s %-20s %-10s\n", row[0], row[1], row[2], row[3])
			} else {
				dateParts := strings.Split(row[1], "-")[1]
				if dateParts == fmt.Sprintf("%02d", Month) {
					fmt.Printf("%-5s %-15s %-20s %-10s\n", row[0], row[1], row[2], row[3])
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(sumaryCmd)
	sumaryCmd.Flags().IntVarP(&Month, "month", "m", 0, "description of the expense")
}

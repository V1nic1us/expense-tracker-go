package cmd

import (
	"fmt"
	"os"
	"time"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
)

var Description string
var Amount float32
const FileName = "Book1.xlsx"

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new expense",
	Long:  `add a new expense to the list of expenses`,

	Run: func(cmd *cobra.Command, args []string) {
		var f *excelize.File
		var err error

		if _, err := os.Stat(FileName); err == nil {
			f, err = excelize.OpenFile(FileName)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			f = excelize.NewFile()
		}

		defer func() {
			if err := f.Close(); err != nil {
				fmt.Println(err)
			}
		}()

		rows, err := f.GetRows("Sheet1")
		if err != nil {
			fmt.Println(err)
			return
		}
		rowCount := len(rows)
		if rowCount == 0 {
			f.SetCellValue("Sheet1", "A1", "ID")
			f.SetCellValue("Sheet1", "B1", "Date")
			f.SetCellValue("Sheet1", "C1", "Description")
			f.SetCellValue("Sheet1", "D1", "Amount")
			rowCount = 1
		}

		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", rowCount+1), rowCount)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", rowCount+1), time.Now().Format("02-01-2006"))
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", rowCount+1), Description)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", rowCount+1), Amount)

		// Save spreadsheet by the given path.
		if err := f.SaveAs(FileName); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&Description, "description", "d", "", "description of the expense")
	addCmd.Flags().Float32VarP(&Amount, "amount", "a", 0, "amount of the expense")
	addCmd.MarkFlagRequired("description")
	addCmd.MarkFlagRequired("amount")
}

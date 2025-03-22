package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
)

var Id int

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete an expense",
	Long:  `delete an expense from the list of expenses`,
	
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
		for idx, row := range rows {
			if len(row) > 0 && row[0] == fmt.Sprintf("%d", Id) {
				f.RemoveRow("Sheet1", idx+1)
				if err := f.SaveAs(FileName); err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Expense deleted successfully.")
				}
				return
			}
		}
		fmt.Println("Expense not found.")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().IntVarP(&Id, "id", "i", 0, "id of the expense to delete")
	deleteCmd.MarkFlagRequired("id")
}

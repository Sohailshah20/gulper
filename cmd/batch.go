package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sohailshah20/csvbatch/csv"
	"github.com/sohailshah20/csvbatch/db"
	"github.com/sohailshah20/csvbatch/textinput"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(batch)
}

var batch = &cobra.Command{
	Use:   "import",
	Short: "batch import csv files into your database",
	Long:  ".",

	Run: func(cmd *cobra.Command, args []string) {
		options := Options{
			FieldName: &textinput.Output{},
		}
		tprogeam := tea.NewProgram(textinput.InitialModel(options.FieldName, "Enter the path to the csv file"))
		_, err := tprogeam.Run()
		if err != nil {
			cobra.CheckErr(err)
		}

		filePath := options.FieldName.Output
		col, data, err := csv.ReadFile(filePath)
		if err != nil {
			cobra.CheckErr(err)
		}
		db, err := db.NewDb()
		if err != nil {
			cobra.CheckErr(err)
		}

		db.BatchInsert(col, data)
	},
}

type Options struct {
	FieldName *textinput.Output
	FiledType string
}

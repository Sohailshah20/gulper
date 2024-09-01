package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sohailshah20/csvbatch/textinput"
	// "github.com/sohailshah20/csvbatch/csv"
	// "github.com/sohailshah20/csvbatch/db"
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

		qs := []string{"path to the csv file", "enter database connection url"}
		questions := textinput.NewQuestions(qs)
		model := textinput.NewMain(questions)
		tprogeam := tea.NewProgram(model, tea.WithAltScreen())
		_, err := tprogeam.Run()
		if err != nil {
			fmt.Println(err.Error())
		}
		if err != nil {
			cobra.CheckErr(err)
		}
		fmt.Println(model.Questions)
	},
}

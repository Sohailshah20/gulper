/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	// "github.com/charmbracelet/lipgloss"

	// "github.com/sohailshah20/csvbatch/cmd"
	"fmt"

	"github.com/sohailshah20/csvbatch/textinput"
	tea "github.com/charmbracelet/bubbletea"
)


func main() {
	// cmd.Execute()

	qs := []string{"path to the csv file", "enter database connection url"}
	questions := textinput.NewQuestions(qs)
	model := textinput.NewMain(questions)
	tprogeam := tea.NewProgram(model, tea.WithAltScreen())
	_, err := tprogeam.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(model.Questions)

}

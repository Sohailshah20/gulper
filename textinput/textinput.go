package textinput

// A simple program demonstrating the text input component from the Bubbles
// component library.

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Output struct {
	Output string
}

func (o *Output) updateOutput(val string) {
	o.Output = val
}

type (
	errMsg error
)

type model struct {
	textInput textinput.Model
	err       error
	output    *Output
	header    string
}

func InitialModel(output *Output, header string) model {
	ti := textinput.New()
	ti.Placeholder = "file path"
	ti.Focus()
	ti.CharLimit = 1024
	ti.Width = 20

	return model{
		textInput: ti,
		err:       nil,
		output:    output,
		header:    header,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.output.updateOutput(m.textInput.Value())
			return m, tea.Quit
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		m.header+"\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

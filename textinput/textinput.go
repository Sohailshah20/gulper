package textinput

// A simple program demonstrating the text input component from the Bubbles
// component library.

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	// "github.com/sohailshah20/csvbatch/cmd"
	// "golang.org/x/text/cases"
)

var style = lipgloss.NewStyle().
	Width(100).
	Align(lipgloss.Left).
	Foreground(lipgloss.Color("63")).
	BorderStyle(lipgloss.RoundedBorder()).
    BorderForeground(lipgloss.Color("228")).
    BorderBackground(lipgloss.Color("63")).
    BorderTop(true).
    BorderLeft(true)


type Output struct {
	Output string
}

func (o *Output) updateOutput(val string) {
	o.Output = val
}

type (
	errMsg error
)

type Question struct {
	question string
	answer   string
}

func NewQuestions(q []string) []Question {
	var questions []Question
	for _, qes := range q {
		questions = append(questions, Question{
			question: qes,
		})
	}
	return questions
}

type Main struct {
	index     int
	Questions []Question
	width     int
	height    int
	answer    textinput.Model
	done      bool
	style     *lipgloss.Style
}

func NewMain(questions []Question) *Main {
	textInput := textinput.New()
	textInput.Placeholder = "type here"
	textInput.Focus()
	return &Main{
		Questions: questions,
		answer:    textInput,
	}
}

func (m Main) Init() tea.Cmd {
	return textinput.Blink
}

func (m *Main) Next() {
	if m.index < len(m.Questions)-1 {
		m.index++
	} else {
		m.index = 0
	}
}

func (m Main) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	current := &m.Questions[m.index]
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.index == len(m.Questions)-1 {
				m.done = true
			}
			current.answer = m.answer.Value()
			m.answer.SetValue("")
			m.Next()
			return m, nil
		}
	}
	m.answer, cmd = m.answer.Update(msg)
	return m, cmd
}

func (m Main) View() string {
	if m.done {
		return "Working"
	}
	if m.width == 0 {
		return "loading..."
	}
	return lipgloss.JoinVertical(
		lipgloss.Center,
		style.Render(m.Questions[m.index].question),
		style.Render(m.answer.View()),
	)
}

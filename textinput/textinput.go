package textinput

// A simple program demonstrating the text input component from the Bubbles
// component library.

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sohailshah20/csvbatch/csv"
	"github.com/sohailshah20/csvbatch/db"
)

type (
	errMsg error
)

type Question struct {
	question string
	answer   string
	done     bool
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
	inserted  bool
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
				return m, nil
			}
			if m.inserted {
				return m, tea.Quit
			}
			current.answer = m.answer.Value()
			current.done = true
			m.answer.SetValue("")
			m.Next()
			return m, nil
		}
	}
	m.answer, cmd = m.answer.Update(msg)
	return m, cmd
}

func (m Main) View() string {
	done := strings.Builder{}
	for _, s := range m.Questions {
		if s.done {
			done.WriteString(s.question + "\n" + s.answer + "\n")
		}
	}
	content := done.String()
	if m.done {
		filePath := m.Questions[0].answer
		col, data, err := csv.ReadFile(filePath)
		if err != nil {
			return err.Error()
		}
		db, err := db.NewDb(m.Questions[1].answer)
		if err != nil {
			return err.Error()
		}

		db.BatchInsert(col, data)
		m.inserted = true
		return "done"
	}
	if m.width == 0 {
		return "loading..."
	}
	// m.Questions.

	return lipgloss.JoinVertical(
		lipgloss.Left,
		content,
		m.Questions[m.index].question,
		m.answer.View(),
	)
}

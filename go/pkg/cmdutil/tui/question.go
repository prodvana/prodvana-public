package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Question struct {
	question  string
	input     textinput.Model
	submitted bool
	response  string
	secret    bool
}

func NewQuestion(question string, placeholder string, secret bool) *Question {
	ti := textinput.New()
	ti.Placeholder = placeholder
	if secret {
		ti.EchoMode = textinput.EchoPassword
	}
	ti.Focus()
	return &Question{
		question: question,
		input:    ti,
		secret:   secret,
	}
}

func (q Question) Init() tea.Cmd {
	return textinput.Blink
}

func (q *Question) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			q.submitted = true
			q.response = strings.TrimSpace(q.input.Value())
			return q, tea.Quit
		case tea.KeyEsc, tea.KeyCtrlC:
			return q, tea.Quit
		}
	}

	q.input, cmd = q.input.Update(msg)
	return q, cmd
}

func (q Question) View() string {
	if q.submitted {
		if q.secret {
			return "\n" + q.question + " " + strings.Repeat("*", len(q.response)) + "\n"
		} else {
			return "\n" + q.question + " " + q.response + "\n"
		}
	}
	return "\n" + q.question + "\n" + q.input.View() + "\n"
}

// QuestionInteractive prompts the user for a freeform text response.
func QuestionInteractive(question, placeholder string) (string, error) {
	quest := NewQuestion(question, placeholder, false)
	if _, err := tea.NewProgram(quest).Run(); err != nil {
		return "", err
	}
	return quest.response, nil
}

func QuestionInteractiveSecret(question, placeholder string) (string, error) {
	quest := NewQuestion(question, placeholder, true)
	if _, err := tea.NewProgram(quest).Run(); err != nil {
		return "", err
	}
	return quest.response, nil
}

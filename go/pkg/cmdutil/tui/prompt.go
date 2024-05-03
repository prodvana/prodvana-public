package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Prompt struct {
	question    string
	input       textinput.Model
	submitted   bool
	accepted    bool
	defaultResp bool
}

func NewPrompt(question string, defaultResp bool) *Prompt {
	ti := textinput.New()
	if defaultResp {
		ti.Placeholder = "yes"
	} else {
		ti.Placeholder = "no"
	}
	ti.Focus()
	ti.CharLimit = 3
	ti.Width = 20
	return &Prompt{
		question:    question,
		input:       ti,
		defaultResp: defaultResp,
	}
}

func (p Prompt) Init() tea.Cmd {
	return textinput.Blink
}

func (p *Prompt) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			p.submitted = true
			answer := strings.TrimSpace(strings.ToLower(p.input.Value()))
			if answer == "" {
				p.accepted = p.defaultResp
			} else {
				p.accepted = answer == "y" || answer == "yes"
			}
			return p, tea.Quit
		case tea.KeyEsc, tea.KeyCtrlC:
			return p, tea.Quit
		}
	}

	p.input, cmd = p.input.Update(msg)
	return p, cmd
}

func (p Prompt) View() string {
	if p.submitted {
		var resp string
		if p.accepted {
			resp = "yes"
		} else {
			resp = "no"
		}
		return "\n" + p.question + " " + resp + "\n"
	}
	return "\n" + p.question + "\n" + p.input.View() + "\n"
}

// PromptInteractive prompts the user for a yes/no answer, where y or yes is true.
func PromptInteractive(question string, defaultResp bool) (bool, error) {
	prompt := NewPrompt(question, defaultResp)
	if _, err := tea.NewProgram(prompt).Run(); err != nil {
		return false, err
	}
	return prompt.accepted, nil
}

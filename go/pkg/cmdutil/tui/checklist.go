package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	listStyle         = lipgloss.NewStyle().PaddingLeft(2)
	runningTextStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	spinnerStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	waitingStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	successStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	failedStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	inconclusiveStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("3"))
)

type Checklist struct {
	Items          []ChecklistItem
	runningSpinner spinner.Model
}

type State int

const (
	StateWaiting State = iota
	StateRunning
	StateSuccess
	StateFailure
	StateInconclusive
)

type ChecklistItem struct {
	Message string
	State   State
}

func NewChecklist(items []ChecklistItem) *Checklist {
	s := spinner.New(
		spinner.WithSpinner(spinner.Dot),
		spinner.WithStyle(spinnerStyle),
	)
	return &Checklist{
		Items:          items,
		runningSpinner: s,
	}
}

func (c *Checklist) SetSuccess(index int, message string) {
	c.Items[index].State = StateSuccess
	c.Items[index].Message = message
}

func (c *Checklist) SetFailure(index int, message string) {
	c.Items[index].State = StateFailure
	c.Items[index].Message = message
}

func (c *Checklist) SetInconclusive(index int, message string) {
	c.Items[index].State = StateInconclusive
	c.Items[index].Message = message
}

func (c *Checklist) SetRunning(index int) {
	c.Items[index].State = StateRunning
}

func (c *Checklist) Init() tea.Cmd {
	return c.runningSpinner.Tick
}

func (c *Checklist) Update(msg tea.Msg) (*Checklist, tea.Cmd) {
	var cmd tea.Cmd
	c.runningSpinner, cmd = c.runningSpinner.Update(msg)
	return c, cmd
}

func (c *Checklist) View() string {
	b := strings.Builder{}
	for _, item := range c.Items {
		switch item.State {
		case StateRunning:
			fmt.Fprint(&b, fmt.Sprintf("[ %s] %s", c.runningSpinner.View(), runningTextStyle.Render(item.Message))+"\n")
		case StateSuccess:
			fmt.Fprint(&b, successStyle.Render(fmt.Sprintf("[ ✓ ] %s", item.Message))+"\n")
		case StateFailure:
			fmt.Fprint(&b, failedStyle.Render(fmt.Sprintf("[ x ] %s", item.Message))+"\n")
		case StateInconclusive:
			fmt.Fprint(&b, inconclusiveStyle.Render(fmt.Sprintf("[ ? ] %s", item.Message))+"\n")
		default:
			fmt.Fprint(&b, waitingStyle.Render(fmt.Sprintf("[ . ] %s", item.Message))+"\n")
		}
	}
	return listStyle.Render(b.String())
}

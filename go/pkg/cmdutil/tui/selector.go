package tui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var defaultStyle = lipgloss.NewStyle()

var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("2"))
)

type listItem string

func (i listItem) FilterValue() string { return string(i) }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, li list.Item) {
	i, ok := li.(listItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type Selector struct {
	list     list.Model
	selected string
}

func (s Selector) Init() tea.Cmd {
	return nil
}

func (s *Selector) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return s, tea.Quit
		case "enter":
			selected, ok := s.list.SelectedItem().(listItem)
			if ok {
				s.selected = string(selected)
				return s, tea.Quit
			}
		}
	}
	var cmd tea.Cmd
	s.list, cmd = s.list.Update(msg)
	return s, cmd
}

func (s Selector) View() string {
	return defaultStyle.Render(s.list.View())
}

func GetInteractiveSelection(title string, items []string) (string, error) {
	listItems := make([]list.Item, len(items))
	for idx, item := range items {
		listItems[idx] = listItem(item)
	}

	l := list.New(listItems, itemDelegate{}, 40, 10)
	l.Title = title
	m := &Selector{list: l}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		return "", err
	}
	return m.selected, nil
}

package preview

import (
	"fmt"
	"strings"
	"tui/internal/tui/view/login"
	"tui/internal/tui/view/register"

	tea "charm.land/bubbletea/v2"
)

type Option struct {
	msg string
	callback tea.Cmd
}

type Model struct {
	blink bool
	index int
	options []Option
}

func New() Model {
	return Model{
		options: []Option{
			Option{
				msg: "Login", 
				callback: login.ChangeToCmd(),
			},
			Option{
				msg: "Register", 
				callback: register.ChangeToCmd(),
			},
		},
	}
}

func (m Model) Init() tea.Cmd {
	return TickCmd()
}

func (m Model) Update(msg tea.Msg) (Model,tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Key().Code {
		case tea.KeyUp:
			m.index = max(m.index - 1, 0)
		case tea.KeyDown:
			m.index = min(m.index + 1, len(m.options)-1)
		case tea.KeyEnter:
			return m, m.options[m.index].callback
		}
	case TickMsg:
		m.blink = !m.blink
		return m, TickCmd()
	}
	return m, nil
}

func (m Model) View() tea.View {
	options := []string{}
	for i := range m.options {
		option := m.options[i].msg
		if i == m.index && m.blink {
			option = " > " + option
		} else {
			option = "   " + option
		}
		options = append(options,option)
	}
	s := fmt.Sprintf("Select an Action:\n%s",
		strings.Join(options, "\n"),
	)
	return tea.NewView(s)
}


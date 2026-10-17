package login

import (
	"fmt"
	"strings"
	"tui/internal/tui/components/text"
	"tui/internal/tui/view/home"

	tea "charm.land/bubbletea/v2"
)

type Model struct {
	blink bool
	index int
	options []text.Model
}

func New() Model {
	return Model{
		options: []text.Model{
			text.New("username", true),
			text.New("password", false),
		},
	}
}

func (m *Model) Init() tea.Cmd {
	return TickCmd()
}

func (m Model) Update(msg tea.Msg) (Model,tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Key().Code {
		case tea.KeyUp:
			m.index = max(m.index - 1, 0)
			return m, nil
		case tea.KeyDown:
			m.index = min(m.index + 1, len(m.options) - 1)
			return m, nil
		case tea.KeyEnter:
			username := m.options[0].Text()
			password := m.options[1].Text()
			return m, home.ChangeToCmd(username, password, ChangeToCmd())
		}
		m.options[m.index], cmd = m.options[m.index].Update(msg)
	case TickMsg:
		m.blink = !m.blink
		for i := range m.options {
			if i == m.index {
				m.options[i].Blink(m.blink)
			} else {
				m.options[i].Blink(false)
			}
		}
		return m, TickCmd()
	}
	return m, cmd
}

func (m Model) View() tea.View {
	options := []string{}
	for i := range m.options {
		options = append(options, " " + m.options[i].View().Content)
	}
	s := fmt.Sprintf("Log in:\n%s",
		strings.Join(options, "\n"),
	)
	return tea.NewView(s)
}



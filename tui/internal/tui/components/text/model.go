package text

import (
	"strings"
	"tui/internal/tui/shared/ansi"
	"unicode"

	tea "charm.land/bubbletea/v2"
)

type Model struct {
	blink bool
	index int
	text []string
	field string
	show bool
}

func New(field string, show bool) Model {
	return Model{
		field: field,
		show: show,
	}
}

func (m *Model) Blink(blink bool) {
	m.blink = blink
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model,tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Key().Code {
		case tea.KeyBackspace:
			if len(m.text) > 0 {
				m.text = m.text[:len(m.text)-1]
			}
			return m, nil
		}
		c := msg.String()
		if isValid(c) && len(m.text) < 20 {
			m.text = append(m.text, c)
		}
	}
	return m, nil
}

func isValid(s string) bool {
	if len(s) == 1 {
		for _,c := range s {
			return unicode.IsDigit(c) || unicode.IsLetter(c) || unicode.IsPunct(c) || unicode.IsSymbol(c)
		}
	}
	return false
}

func (m Model) View() tea.View {
	text := strings.Join(m.text, "")
	if !m.show {
		text = strings.Repeat("*", len(m.text))
	}
	s := m.field + ": " + text
	if m.blink {
		s = s + ansi.LIGHT.Print(" ")
	}
	return tea.NewView(s)
}

func (m *Model) Text() string {
	return strings.Join(m.text, "")
}




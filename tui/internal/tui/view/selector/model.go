package selector

import (
	"fmt"
	"strings"
	"tui/internal/client"
	"tui/internal/tui/shared"
	"tui/internal/tui/shared/ansi"

	tea "charm.land/bubbletea/v2"
)

type Credentials client.CredentialsRequest
type Size client.Size

type Model struct {
	ready bool
	blink bool
	index int
	credentials Credentials
	sizes []Size
}

func Header() string {
	return fmt.Sprintf("%-10s  %-5s  %-8s  %-8s  %-8s  %-8s",
		"Specs",
		"CPU",
		"MEM(MB)",
		"DISK(GB)",
		"$/Month",
		"$/Hour",
	)
}

func (s *Size) Display() string {
	return fmt.Sprintf("%-10s  %-5d  %-8d  %-8d  %-8.2f  %-8.2f",
		s.Slug,
		s.VCpu,
		s.Mem,
		s.Disk,
		s.PriceMonth,
		s.PriceHour,
	)
}

func New() Model {
	return Model{
		sizes: []Size{},
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Sequence(ReadyCmd(false), TickCmd())
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Key().Code {
		case tea.KeyEnter:
			//return m, CreateVpsCmd(m.credentials, m.sizes[m.index])
			return m, nil
		case tea.KeyUp:
			m.index = max(0,m.index - 1)
			return m, nil
		case tea.KeyDown:
			m.index = min(len(m.sizes)-1, m.index + 1)
			return m, nil
		}
	case ReadyMsg:
		m.ready = bool(msg)
		return m, nil
	case TickMsg:
		m.blink = !m.blink
		return m, TickCmd()
	case SelectSuccess:
		m.credentials = msg.cred
		m.sizes = msg.sizes
		return m, tea.Batch(
			ReadyCmd(true),
			shared.SendStatusCmd( "Operation Succeded", false ),
		)
	case SelectFailure:

		return m, tea.Batch(
			shared.ChangeViewCmd(shared.LoginView),
			shared.SendStatusCmd( msg.err, true ),
		)
	}
	return m, nil
}

func (m Model) View() tea.View {
	if !m.ready { return tea.NewView("") }
	s := []string{}
	s = append(s, Header())
	for i := range m.sizes {
		d := m.sizes[i].Display()
		if m.index == i && m.blink {
			d = ansi.LIGHT.Print(d)
		}
		s = append(s,d)
	}
	return tea.NewView(strings.Join(s, "\n"))
}

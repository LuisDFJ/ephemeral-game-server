package home

import (
	"fmt"
	"tui/internal/tui/shared"
	"tui/internal/tui/view/selector"

	tea "charm.land/bubbletea/v2"
)

type Client struct {}

type Info struct {
	alive bool
	slug string
	ip string
	world string
	password string
} 

type Credentials struct {
	username string
	password string
}

type Opion struct {
	text string
	callback func(Credentials) tea.Cmd
}

type Model struct {
	ready bool
	blink bool
	index int
	info Info
	client Client
	credentials Credentials
	options []Opion
}

func New() Model {
	return Model{
		options: []Opion{
			Opion{
				text: "Deploy",
				callback: func(c Credentials) tea.Cmd {return selector.ChangeToCmd()},
			},
			Opion{
				text: "Destroy",
				callback: func(c Credentials) tea.Cmd {return ChangeToCmd(c.username, c.password, shared.ChangeViewCmd(shared.HomeView))},
			},
		},
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
			return m, m.options[m.index].callback(m.credentials)
		}
	case ReadyMsg:
		m.ready = bool(msg)
		return m, nil
	case TickMsg:
		m.blink = !m.blink
		return m, TickCmd()
	case LogInSuccessMsg:
		m.info = msg.info
		m.credentials = msg.cred
		if m.info.alive {
			m.index = 1
		} else {
			m.index = 0
		}
		return m, tea.Batch(
			ReadyCmd(true),
			shared.SendStatusCmd("Succesful Login", false),
		)
	case LogInFailureMsg:
		return m, tea.Batch(
			msg.fallback,
			shared.SendStatusCmd(msg.err, true),
		)
	}
	return m, nil
}

func (m Model) View() tea.View {
	if !m.ready { return tea.NewView("") }

	status := "Server Status: offline"
	if m.info.alive {
		status = fmt.Sprintf("Server Status: online\nIP: %s\nSpecs: %s\nWorld: %s    Password: %s",
			m.info.ip, m.info.slug, m.info.world, m.info.password,
		)
	}
	option := m.options[m.index].text
	if m.blink {
		option = " > " + option
	} else {
		option = "   " + option
	}
	s := fmt.Sprintf("%s\n%s\n", status, option )
	return tea.NewView(s)
}


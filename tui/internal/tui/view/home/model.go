package home

import (
	"fmt"
	"tui/internal/client"
	"tui/internal/tui/shared"
	"tui/internal/tui/shared/ansi"
	"tui/internal/tui/view/selector"

	tea "charm.land/bubbletea/v2"
)

type Credentials client.CredentialsRequest
type Status client.StatusResponse

type Opion struct {
	text string
	callback func(Credentials) tea.Cmd
}

type Model struct {
	ready bool
	blink bool
	index int
	status Status
	credentials Credentials
	options []Opion
}

func New() Model {
	return Model{
		options: []Opion{
			Opion{
				text: "Deploy",
				callback: func(c Credentials) tea.Cmd {return selector.ChangeToCmd(c.Username, c.Password, ChangeToCmd(c.Username,c.Password,shared.ChangeViewCmd(shared.LoginView)))},
			},
			Opion{
				text: "Destroy",
				callback: DestroyCallback,
			},
		},
	}
}

func (m Model) Init() tea.Cmd {
	return ReadyCmd(false)
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if !m.ready { return m,nil }
		switch msg.Key().Code {
		case tea.KeyEnter:
			return m, m.options[m.index].callback(m.credentials)
		}
	case ReadyMsg:
		m.ready = bool(msg)
		return m, nil
	case shared.TickMsg:
		m.blink = !m.blink
		return m, nil
	case LogInSuccessMsg:
		m.status = msg.status
		m.credentials = msg.cred
		if m.status.Vps.Alive {
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

	status := "Server Status: " + ansi.ERROR.Print("offline")
	if m.status.Vps.Alive {
		status = fmt.Sprintf("Server Status: %s\nIP: %s\nSpecs: %s\nWorld: %s    Password: %s",
			ansi.INFO.Print("online"),
			ansi.INFO.Print(m.status.Vps.Ip),
			ansi.INFO.Print(m.status.Vps.Specs),
			ansi.INFO.Print(m.status.World.Name),
			ansi.INFO.Print(m.status.World.Password),
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


package home

import (
	"time"
	"tui/internal/tui/shared"

	tea "charm.land/bubbletea/v2"
)

type TickMsg struct {}
func TickCmd() tea.Cmd {
	return tea.Tick(shared.BlinkDuration, func(t time.Time) tea.Msg {return TickMsg{}})
}

type ReadyMsg bool
func ReadyCmd(ready bool) tea.Cmd {
	return func() tea.Msg { return ReadyMsg(ready) }
}

type LogInSuccessMsg struct {
	info Info
	cred Credentials
}

type LogInFailureMsg struct {
	err string
	fallback tea.Cmd
}

func SetupCmd(username string, password string, fallback tea.Cmd) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(time.Second * 2)
		if username == "admin" && password == "password" {
			return LogInSuccessMsg{
				info: Info{alive: true, slug: "c-2", ip: "1.1.1.1", world: "World", password: "password"},
				cred: Credentials{username: username, password: password},
			}
		}
		return LogInFailureMsg{err: "Log In Failed", fallback: fallback}
	}
}

func ChangeToCmd(username string, password string, fallback tea.Cmd) tea.Cmd {
	return tea.Sequence(
		shared.ChangeViewCmd(shared.HomeView),
		shared.LoadableCmd(SetupCmd(username,password, fallback)),
	)
}

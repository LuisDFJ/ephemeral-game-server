package home

import (
	"time"
	"tui/internal/client"
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
	status Status
	cred Credentials
}

type LogInFailureMsg struct {
	err string
	fallback tea.Cmd
}

func SetupCmd(username string, password string, fallback tea.Cmd) tea.Cmd {
	return func() tea.Msg {
		c, err := client.New(username, password, "http://localhost:8080")
		if err != nil {
			return LogInFailureMsg{err: "Log In Failed: " + err.Error(), fallback: fallback}
		}
		status, err := c.GetStatus()
		if err != nil {
			return LogInFailureMsg{err: "Could not get Status: " + err.Error(), fallback: fallback}
		}
		return LogInSuccessMsg{
			status: Status(status),
			cred: Credentials{Username: username, Password: password},
		}
	}
}

func ChangeToCmd(username string, password string, fallback tea.Cmd) tea.Cmd {
	return tea.Sequence(
		shared.ChangeViewCmd(shared.HomeView),
		shared.LoadableCmd(SetupCmd(username,password, fallback)),
	)
}

func DestroyCmd(username string, password string) tea.Cmd {
	return func() tea.Msg {
		c, err := client.New(username, password, "http://localhost:8080")
		if err != nil {
			return LogInFailureMsg{err: "Authentication Failure: " + err.Error(), fallback: shared.ChangeViewCmd(shared.LoginView)}
		}
		err = c.PostDestroyVps()
		if err != nil {
			return LogInFailureMsg{err: "Could not destroy VPS: " + err.Error(), fallback: shared.ChangeViewCmd(shared.HomeView)}
		}
		time.Sleep(2*time.Second)
		status, err := c.GetStatus()
		if err != nil {
			return LogInFailureMsg{err: "Could not get Status: " + err.Error(), fallback: shared.ChangeViewCmd(shared.LoginView)}
		}
		return LogInSuccessMsg{
			status: Status(status),
			cred: Credentials{Username: username, Password: password},
		}
	}
}

func DestroyCallback(cred Credentials) tea.Cmd {
	return tea.Sequence(
		ReadyCmd(false),
		shared.LoadableCmd(DestroyCmd(cred.Username, cred.Password)),
	)
}

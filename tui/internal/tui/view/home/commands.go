package home

import (
	"time"
	"tui/internal/client"
	"tui/internal/tui/shared"

	tea "charm.land/bubbletea/v2"
)

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
		c, err := client.New(username, password)
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
		c, err := client.New(username, password)
		if err != nil {
			return LogInFailureMsg{err: "Authentication Failure: " + err.Error(), fallback: shared.ChangeViewCmd(shared.LoginView)}
		}
		err = c.PostDestroyVps()
		if err != nil {
			return LogInFailureMsg{err: "Could not destroy VPS: " + err.Error(), fallback: shared.ChangeViewCmd(shared.HomeView)}
		}
		time.Sleep(10*time.Second)
		return shared.SendStatusCmd("Succesfully destroyed VPS", false)
	}
}

func DestroyCallback(cred Credentials) tea.Cmd {
	return tea.Sequence(
		ReadyCmd(false),
		shared.LoadableCmd(DestroyCmd(cred.Username, cred.Password)),
		ChangeToCmd(cred.Username, cred.Password, shared.ChangeViewCmd(shared.LoginView)),
	)
}

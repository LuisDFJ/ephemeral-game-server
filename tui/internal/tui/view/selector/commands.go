package selector

import (
	"tui/internal/client"
	"tui/internal/tui/shared"

	tea "charm.land/bubbletea/v2"
)

type ReadyMsg bool
func ReadyCmd(ready bool) tea.Cmd {
	return func() tea.Msg { return ReadyMsg(ready) }
}

type TickMsg struct {}
func TickCmd() tea.Cmd {
	return func() tea.Msg { return TickMsg{} }
}

type SelectSuccess struct {
	cred Credentials
	sizes []Size
}

type SelectFailure struct {
	err string
}

func GetSizesCmd(user string, password string) tea.Cmd {
	return func() tea.Msg {
		c,err := client.New(user,password,"http://localhost:8080")
		if err != nil {
			return SelectFailure{ err: "Failed to connect to server: " + err.Error() }
		}
		sizes,err := c.GetSizes()
		if err != nil {
			return SelectFailure{ err: "Failed to retrieve sizes: " + err.Error() }
		}
		if len(sizes) == 0 {
			return SelectFailure{ err: "Servers unavailable for now retry later" }
		}
		res := []Size{}
		for _,s := range sizes {
			res = append(res, Size(s))
		}
		return SelectSuccess{
			sizes: res,
			cred: Credentials{Username: user, Password: password},
		}
	}
}

func ChangeToCmd(user string, password string) tea.Cmd {
	return tea.Sequence(
		shared.ChangeViewCmd(shared.SelectView),
		shared.LoadableCmd(GetSizesCmd(user,password)),
	)
}

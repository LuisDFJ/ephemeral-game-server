package selector

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

type SelectSuccess struct {
	cred Credentials
	sizes []Size
	fallback tea.Cmd
}

type SelectFailure struct {
	err string
	fallback tea.Cmd
}

func GetSizesCmd(user string, password string, fallback tea.Cmd) tea.Cmd {
	return func() tea.Msg {
		c,err := client.New(user,password)
		if err != nil {
			return SelectFailure{ err: "Failed to connect to server: " + err.Error(), fallback: fallback }
		}
		sizes,err := c.GetSizes()
		if err != nil {
			return SelectFailure{ err: "Failed to retrieve sizes: " + err.Error(), fallback: fallback }
		}
		if len(sizes) == 0 {
			return SelectFailure{ err: "Servers unavailable for now retry later", fallback: fallback }
		}
		res := []Size{}
		for _,s := range sizes {
			res = append(res, Size(s))
		}
		return SelectSuccess{
			sizes: res,
			cred: Credentials{Username: user, Password: password},
			fallback: fallback,
		}
	}
}

func ChangeToCmd(user string, password string, fallback tea.Cmd) tea.Cmd {
	return tea.Sequence(
		shared.ChangeViewCmd(shared.SelectView),
		shared.LoadableCmd(GetSizesCmd(user,password, fallback)),
	)
}

type VpsCreationSuccessMsg struct {

}

func createVps(user string, password string, slug string) tea.Cmd {
	return func() tea.Msg {
		c,err := client.New(user,password)
		if err != nil {
			return SelectFailure{ err: "Failed to connect to server: " + err.Error(), fallback: nil }
		}
		err = c.PostCreateVps(slug)
		if err != nil {
			return SelectFailure{ err: "Failed to create vps: " + err.Error(), fallback: nil }
		}

		for _ = range 20 {
			time.Sleep(time.Second * 6)
			s,err := c.GetStatus()
			if err != nil {
				return SelectFailure{ err: "Failed to reach server", fallback: nil }
			}
			if s.Vps.Ip != "" { break }
		}
		return VpsCreationSuccessMsg{}
	}
}

func CreateVpsCmd(user string, password string, size Size, fallback tea.Cmd) tea.Cmd {
	return tea.Sequence(
		shared.LoadableCmd(createVps(user,password,size.Slug)),
		fallback,
	)
}

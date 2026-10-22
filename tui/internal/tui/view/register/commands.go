package register

import (
	"os"
	"tui/internal/client"
	"tui/internal/tui/shared"
	tea "charm.land/bubbletea/v2"
)

func ChangeToCmd() tea.Cmd {
	return shared.ChangeViewCmd(shared.RegisterView)
}

type RegisterFailedMsg struct {
	err string
}
type RegisterSuccedMsg struct {}
func RegisterCmd(user string, password string, confirm string) tea.Cmd {
	return func() tea.Msg {
		if password != confirm {
			return RegisterFailedMsg{ err : "Password Mismatch" }
		}
		ott, err := os.ReadFile("ott.env")
		if err != nil {
			return RegisterFailedMsg{ err : "Missing ott.env file" }
		}
		err = client.Register(string(ott), user, password)
		if err != nil {
			return RegisterFailedMsg{ err : err.Error() }
		}
		os.Remove("ott.env")
		return RegisterSuccedMsg{}
	}
}

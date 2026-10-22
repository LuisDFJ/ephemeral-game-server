package login

import (
	"tui/internal/tui/shared"

	tea "charm.land/bubbletea/v2"
)

func ChangeToCmd() tea.Cmd {
	return shared.ChangeViewCmd(shared.LoginView)
}

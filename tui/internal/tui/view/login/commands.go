package login

import (
	"time"
	"tui/internal/tui/shared"

	tea "charm.land/bubbletea/v2"
)

type TickMsg struct {}
func TickCmd() tea.Cmd {
	return tea.Tick(shared.BlinkDuration, func(t time.Time) tea.Msg {return TickMsg{}})
}

func ChangeToCmd() tea.Cmd {
	return shared.ChangeViewCmd(shared.LoginView)
}

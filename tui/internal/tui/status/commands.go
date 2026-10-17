package status

import (
	"time"

	tea "charm.land/bubbletea/v2"
)

type TickMsg struct {}
func TickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond * 100, func(t time.Time) tea.Msg {return TickMsg{}})
}

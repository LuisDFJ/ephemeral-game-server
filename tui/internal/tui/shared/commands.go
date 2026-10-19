package shared

import (
	"time"

	tea "charm.land/bubbletea/v2"
)

const MsgTimeout = time.Second * 5
const BlinkDuration = time.Millisecond * 400

type View int
const (
	PreView View = iota
	LoginView
	RegisterView
	HomeView
	SelectView
)

type ChangeViewMsg View
func ChangeViewCmd(view View) tea.Cmd {
	return func() tea.Msg {
		return ChangeViewMsg(view)
	}
}

type LoadingMsg bool
func LoadingCmd(load bool) tea.Cmd {
	return func() tea.Msg {
		return LoadingMsg(load)
	}
}

func LoadableCmd(cmd tea.Cmd) tea.Cmd {
	return tea.Sequence( tea.Batch(cmd, LoadingCmd(true)), LoadingCmd(false) )
}

type StatusMsg struct { Msg string; Err bool }
func StatusCmd(msg string, err bool) tea.Cmd {
	return func() tea.Msg {
		return StatusMsg{Msg: msg, Err: err}
	}
}

type ResetStatusMsg struct {}

func SendStatusCmd(msg string, err bool) tea.Cmd {
	cmd := tea.Tick(MsgTimeout, func(t time.Time) tea.Msg { return ResetStatusMsg{} })
	return tea.Batch(StatusCmd(msg,err), cmd)
}


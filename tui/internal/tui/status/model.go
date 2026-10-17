package status

import (
	"fmt"
	"tui/internal/tui/shared"
	"tui/internal/tui/shared/ansi"

	tea "charm.land/bubbletea/v2"
)


type Model struct {
	msg string
	err bool
	loading bool
	sequence []string
	count int
}

func New() Model {
	return Model{
		sequence: []string{"-", "\\", "|", "/"},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model,tea.Cmd) {
	cmds := []tea.Cmd{}
	switch msg := msg.(type) {
	case shared.LoadingMsg:
		m.loading = bool(msg)
		if m.loading { cmds = append(cmds, TickCmd()) }
	case shared.StatusMsg:
		m.msg = msg.Msg
		m.err = msg.Err
	case shared.ResetStatusMsg:
		m.msg = ""
		m.err = false
	case TickMsg:
		m.count = (m.count + 1) % len(m.sequence)
		if m.loading { cmds = append(cmds, TickCmd()) }
	}
	return m, tea.Batch(cmds...)
}

func (m Model) View() tea.View {
	loader := "            "
	if m.loading { loader = m.sequence[m.count] + " Loading..." }
	loader = ansi.LOADER.Print(loader)
	var msg string
	if m.err {
		msg = ansi.ERROR.Print(m.msg)
	} else {
		msg = ansi.INFO.Print(m.msg)
	}

	s := fmt.Sprintf("  %s   %s",
		loader,
		msg,
	)
	return tea.NewView(s)
}



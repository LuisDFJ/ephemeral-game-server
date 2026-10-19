package tui

import (
	"fmt"

	"tui/internal/tui/shared"
	"tui/internal/tui/status"
	"tui/internal/tui/view/home"
	"tui/internal/tui/view/login"
	"tui/internal/tui/view/preview"
	"tui/internal/tui/view/register"
	"tui/internal/tui/view/selector"

	tea "charm.land/bubbletea/v2"
)


type Model struct {
	view shared.View
	status status.Model
	preview preview.Model
	login login.Model
	register register.Model
	home home.Model
	selector selector.Model
}

func New() Model {
	return Model{
		view: shared.PreView,
		status: status.New(),
		preview: preview.New(),
		login: login.New(),
		register: register.New(),
		home: home.New(),
		selector: selector.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return shared.ChangeViewCmd(shared.PreView)
}

func (m Model) Update(msg tea.Msg) (tea.Model,tea.Cmd) {
	cmds := []tea.Cmd{}

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
	}
	case shared.ChangeViewMsg:
		m.view = shared.View(msg)
		return m, m.muxInit()
	}

	var cmd tea.Cmd
	m.status, cmd = m.status.Update(msg)
	cmds = append(cmds, cmd)

	cmds = append(cmds, m.muxMsg(msg))

	return m, tea.Batch(cmds...)
}

func (m Model) View() tea.View {
	var view tea.View = m.muxView()

	s := fmt.Sprintf("\n%s\n%s\n",
		m.status.View().Content,
		view.Content,
	)
	return tea.NewView(s)
}


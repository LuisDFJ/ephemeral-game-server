package tui

import (
	"tui/internal/tui/shared"
	tea "charm.land/bubbletea/v2"
)

func (m *Model) muxInit() tea.Cmd {
	var cmd tea.Cmd
	switch m.view {
	case shared.PreView:
		cmd = m.preview.Init()
	case shared.RegisterView:
		cmd = m.register.Init()
	case shared.LoginView:
		cmd = m.login.Init()
	case shared.HomeView:
		cmd = m.home.Init()
	case shared.SelectView:
	}
	return cmd
}

func (m *Model) muxMsg( msg tea.Msg ) tea.Cmd {
	var cmd tea.Cmd
	switch m.view {
	case shared.PreView:
		m.preview, cmd = m.preview.Update(msg)
	case shared.RegisterView:
		m.register, cmd = m.register.Update(msg)
	case shared.LoginView:
		m.login, cmd = m.login.Update(msg)
	case shared.HomeView:
		m.home, cmd = m.home.Update(msg)
	case shared.SelectView:
	}
	return cmd
}

func (m *Model) muxView() tea.View {
	var view tea.View
	switch m.view {
	case shared.PreView:
		view = m.preview.View()
	case shared.RegisterView:
		view = m.register.View()
	case shared.LoginView:
		view = m.login.View()
	case shared.HomeView:
		view = m.home.View()
	case shared.SelectView:
	}
	return view
}

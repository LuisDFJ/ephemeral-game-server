package main

import (
	"os"
	"fmt"
	"tui/internal/tui"
	tea "charm.land/bubbletea/v2"
)


func main() {
	tui := tea.NewProgram(tui.New())
	if _,err := tui.Run(); err != nil {
		fmt.Printf("Failed to run TUI: %v\n", err)
		os.Exit(1)
	}
}

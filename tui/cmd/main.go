package main

import (
	"fmt"
	"os"
	"tui/internal/tui"

	tea "charm.land/bubbletea/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Unable to find .env file")
		os.Exit(1)
	}
	tui := tea.NewProgram(tui.New())
	if _,err := tui.Run(); err != nil {
		fmt.Printf("Failed to run TUI: %v\n", err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dbarrosop/ttoys/ui"
	"github.com/urfave/cli/v2"
)

var Version string

func main() {
	app := &cli.App{
		Name:    "ttoys",
		Usage:   "ttoys is a tool to help you with your development",
		Version: Version,
		Action: func(*cli.Context) error {
			p := tea.NewProgram(ui.New(), tea.WithAltScreen())
			if _, err := p.Run(); err != nil {
				return fmt.Errorf("Alas, there's been an error: %w", err)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

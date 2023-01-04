package style

import "github.com/charmbracelet/lipgloss"

var (
	ColorBlue   = lipgloss.AdaptiveColor{Light: "#ace4ef", Dark: "#ace4ef"}
	ColorGray   = lipgloss.AdaptiveColor{Light: "#999", Dark: "#666"}
	ColorRed    = lipgloss.AdaptiveColor{Light: "#f00", Dark: "#f00"}
	ColorYellow = lipgloss.AdaptiveColor{Light: "#f0c674", Dark: "#f0c674"}
)

var (
	styleTitle = lipgloss.NewStyle().
			Foreground(ColorYellow).
			Bold(true).
			Render

	styleSection = lipgloss.NewStyle().
			Foreground(ColorBlue).
			Bold(true).
			Render

	styleHelp = lipgloss.NewStyle().
			Foreground(ColorGray).
			Render

	styleError = lipgloss.NewStyle().
			Foreground(ColorRed).
			Render
)

func Title(s string) string {
	return styleTitle(s)
}

func Section(s string) string {
	return styleSection(s)
}

func Error(s string) string {
	return styleError(s)
}

func Help(s string) string {
	return styleHelp(s)
}

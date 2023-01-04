package component

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dbarrosop/ttoys/ui/style"
)

type Component struct {
	help   help.Model
	Width  int
	Length int
}

func New() *Component {
	return &Component{
		help: help.New(),
	}
}

func (c *Component) OnKeyPress(
	msg tea.KeyMsg, m tea.Model, parent tea.Model, inputList *InputList,
) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, KeyQuit):
		return m, tea.Quit
	case key.Matches(msg, KeyBack):
		if parent != nil {
			return parent, nil
		}
		return m, tea.Quit
	case key.Matches(msg, KeyFocusNext):
		inputList.Next()
	case key.Matches(msg, KeyFocusPrev):
		inputList.Prev()
	case key.Matches(msg, KeyAcceptInput):
		inputList.Accept()
	case key.Matches(msg, KeyResetInput):
		inputList.Reset()
	case key.Matches(msg, KeyHelp):
		c.help.ShowAll = !c.help.ShowAll
	}
	return m, nil
}

func (c *Component) OnWindowSizeMsg(msg tea.WindowSizeMsg) {
	c.Width = msg.Width
	c.Length = msg.Height

	c.help.Width = msg.Width
}

func (c *Component) Separator() string {
	return strings.Repeat("─", c.Width)
}

func (c *Component) DotSeparator() string {
	return strings.Repeat("·", c.Width)
}

func (c *Component) AvailableScreen(bindings [][]key.Binding) int {
	helpView := c.help.View(
		&KeyViewHelp{
			Long: bindings,
		},
	)

	return c.Length - 3 - len(strings.Split(helpView, "\n"))
}

func (c *Component) RenderView(title, body string, bindings [][]key.Binding) string {
	s := style.Title(title + "\n" + c.Separator())
	s += "\n"

	s += body

	s += "\n"

	helpView := c.help.View(
		&KeyViewHelp{
			Long: bindings,
		},
	)

	padding := c.Length - len(strings.Split(s, "\n")) - len(strings.Split(helpView, "\n"))
	if padding > 0 {
		s += strings.Repeat("\n", padding)
	}

	s += helpView

	return s
}

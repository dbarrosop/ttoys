package component

import "github.com/charmbracelet/bubbles/key"

var (
	KeyAcceptInput = key.NewBinding(
		key.WithKeys("alt+enter"),
		key.WithHelp("ALT+ENTER", "Accept input"),
	)
	KeyBack = key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("ESC", "Back"),
	)
	KeyFocusNext = key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("TAB", "Next input"),
	)
	KeyFocusPrev = key.NewBinding(
		key.WithKeys("shift+tab"),
		key.WithHelp("SHIFT+TAB", "Prev input"),
	)
	KeyHelp = key.NewBinding(
		key.WithKeys("f1"),
		key.WithHelp("F1", "Help"),
	)
	KeyResetInput = key.NewBinding(
		key.WithKeys("ctrl+d"),
		key.WithHelp("CTRL+D", "Reset input"),
	)
	KeyQuit = key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("CTRL+C", "Quit"),
	)
)

type KeyViewHelp struct {
	Long [][]key.Binding
}

func (k *KeyViewHelp) ShortHelp() []key.Binding {
	return []key.Binding{KeyHelp, KeyBack, KeyQuit}
}

func (k *KeyViewHelp) FullHelp() [][]key.Binding {
	r := make([][]key.Binding, len(k.Long)+2)
	r[0] = []key.Binding{KeyHelp, KeyBack, KeyQuit}
	r[1] = []key.Binding{KeyFocusNext, KeyFocusPrev, KeyAcceptInput, KeyResetInput}
	copy(r[2:], k.Long)
	return r
}

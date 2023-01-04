package uuid

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dbarrosop/ttoys/ui/component"
	"github.com/google/uuid"
)

type Model struct {
	component *component.Component
	parent    tea.Model

	generated []string
}

func New(comp *component.Component, parent tea.Model) *Model {
	return &Model{
		component: comp,
		parent:    parent,
		generated: make([]string, 10),
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.component.OnWindowSizeMsg(msg)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, component.KeyAcceptInput):
			for i := 0; i < 10; i++ {
				m.generated[i] = uuid.New().String()
			}
			return m, nil
		default:
			return m.component.OnKeyPress(msg, m, m.parent, nil)
		}
	}

	return m, nil
}

func (m *Model) View() string {
	if m.generated[0] == "" {
		for i := 0; i < 10; i++ {
			m.generated[i] = uuid.New().String()
		}
	}

	var s string
	for _, u := range m.generated {
		s += u
		s += "\n"
	}

	return m.component.RenderView("UUID Generator", s, [][]key.Binding{})
}

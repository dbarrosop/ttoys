package ui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dbarrosop/ttoys/ui/component"
	"github.com/dbarrosop/ttoys/ui/model/base64"
	"github.com/dbarrosop/ttoys/ui/model/date"
	"github.com/dbarrosop/ttoys/ui/model/json"
	"github.com/dbarrosop/ttoys/ui/model/uuid"
)

type keyMap struct {
	One   key.Binding
	Two   key.Binding
	Three key.Binding
	Four  key.Binding
}

func (k *keyMap) Bindings() [][]key.Binding {
	return [][]key.Binding{
		{k.One, k.Two},
	}
}

type UI struct {
	keyMap keyMap

	component *component.Component

	uuidModelFunc   func(*component.Component, tea.Model) *uuid.Model
	jsonModelFunc   func(*component.Component, tea.Model) *json.Model
	base64ModelFunc func(*component.Component, tea.Model) *base64.Model
	dateModelFunc   func(*component.Component, tea.Model) *date.Model
}

func New() *UI {
	ui := &UI{
		component: component.New(),

		uuidModelFunc:   uuid.New,
		jsonModelFunc:   json.New,
		base64ModelFunc: base64.New,
		dateModelFunc:   date.New,

		keyMap: keyMap{
			One: key.NewBinding(
				key.WithKeys("1"),
			),
			Two: key.NewBinding(
				key.WithKeys("2"),
			),
			Three: key.NewBinding(
				key.WithKeys("3"),
			),
			Four: key.NewBinding(
				key.WithKeys("4"),
			),
		},
	}

	return ui
}

func (m *UI) Init() tea.Cmd {
	return nil
}

func (m *UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.component.OnWindowSizeMsg(msg)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.One):
			return m.uuidModelFunc(m.component, m), nil
		case key.Matches(msg, m.keyMap.Two):
			return m.jsonModelFunc(m.component, m), nil
		case key.Matches(msg, m.keyMap.Three):
			return m.base64ModelFunc(m.component, m), nil
		case key.Matches(msg, m.keyMap.Four):
			return m.dateModelFunc(m.component, m), nil
		default:
			return m.component.OnKeyPress(msg, m, nil, nil)
		}
	}

	return m, nil
}

func (m *UI) View() string {
	s := "1. UUID Generator\n"
	s += "2. JSON Viewer\n"
	s += "3. Base64 Convertor\n"
	s += "4. Date Formatter\n"

	return m.component.RenderView(
		"Developer Tools (Terminal Edition)", s, m.keyMap.Bindings(),
	)
}

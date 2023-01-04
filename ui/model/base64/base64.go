package base64

import (
	"encoding/base64"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dbarrosop/ttoys/ui/component"
	"github.com/dbarrosop/ttoys/ui/style"
)

func encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func decode(s string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

type Model struct {
	component *component.Component
	parent    tea.Model

	inputList *component.InputList
	encTa     textarea.Model
	decTa     textarea.Model

	err error
}

func New(comp *component.Component, parent tea.Model) *Model {
	encTa := textarea.New()
	encTa.Prompt = ""
	encTa.ShowLineNumbers = false
	encTa.SetWidth(comp.Width)
	encTa.Focus()
	encTa.CharLimit = -1

	decTa := textarea.New()
	decTa.Prompt = ""
	decTa.ShowLineNumbers = false
	decTa.SetWidth(comp.Width)
	decTa.Blur()
	decTa.CharLimit = -1

	m := &Model{
		component: comp,
		parent:    parent,
		encTa:     encTa,
		decTa:     decTa,
		inputList: component.NewInputList(),
	}

	m.inputList.Add(
		component.Input{
			Input: &m.encTa,
			AcceptAction: func() {
				var v string
				v, m.err = decode(m.encTa.Value())

				if m.err == nil {
					m.decTa.SetValue(v)
					m.decTa.SetCursor(len(v))
				}
			},
		},
	)

	m.inputList.Add(
		component.Input{
			Input: &m.decTa,
			AcceptAction: func() {
				v := encode(m.decTa.Value())
				m.err = nil

				m.encTa.SetValue(v)
				m.encTa.SetCursor(len(v))
			},
		},
	)

	return m
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	tiCmd := m.inputList.Update(msg)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.component.OnWindowSizeMsg(msg)

		m.encTa.SetWidth(msg.Width)
		m.decTa.SetWidth(msg.Width)

	case tea.KeyMsg:
		return m.component.OnKeyPress(msg, m, m.parent, m.inputList)
	}

	return m, tiCmd
}

func (m *Model) View() string {
	bindings := [][]key.Binding{}

	screenSize := m.component.AvailableScreen(bindings)

	m.encTa.SetHeight(screenSize/2 - 2)
	m.decTa.SetHeight(screenSize/2 - 2)

	s := style.Section("Encoded:") + "\n"
	s += m.encTa.View() + "\n\n"
	s += style.Section("Decoded:") + "\n"
	s += m.decTa.View() + "\n"

	if m.err != nil {
		s += style.Error(m.err.Error())
	}

	return m.component.RenderView("Base64 Convertor", s, bindings)
}

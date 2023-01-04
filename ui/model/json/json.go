package json

import (
	"encoding/json"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dbarrosop/ttoys/ui/component"
	"github.com/dbarrosop/ttoys/ui/style"
	"github.com/itchyny/gojq"
)

type Model struct {
	component *component.Component
	parent    tea.Model

	inputList *component.InputList
	jsonTa    textarea.Model
	jqTa      textarea.Model

	json  json.RawMessage
	query *gojq.Query

	err error
}

func New(comp *component.Component, parent tea.Model) *Model {
	jsonTa := textarea.New()
	jsonTa.Focus()
	jsonTa.CharLimit = -1
	jsonTa.SetHeight(10)
	jsonTa.SetWidth(comp.Width)
	jsonTa.SetValue(`{"products":[{"id":1},{"id":2}]}`)

	jqTa := textarea.New()
	jqTa.SetHeight(3)
	jqTa.SetWidth(comp.Width)
	jqTa.CharLimit = -1
	jqTa.SetValue(".products[].id")

	m := &Model{
		component: comp,
		parent:    parent,
		jsonTa:    jsonTa,
		jqTa:      jqTa,
		inputList: component.NewInputList(),
	}

	m.inputList.Add(component.Input{
		Input: &m.jsonTa,
		AcceptAction: func() {
			m.err = json.Unmarshal([]byte(m.jsonTa.Value()), &m.json)

			if m.jqTa.Value() != "" {
				m.query, m.err = gojq.Parse(m.jqTa.Value())
			} else {
				m.query = nil
			}
		},
	})
	m.inputList.Add(component.Input{
		Input: &m.jqTa,
		AcceptAction: func() {
			if m.jqTa.Value() != "" {
				m.query, m.err = gojq.Parse(m.jqTa.Value())
			} else {
				m.query = nil
			}
		},
	})

	return m
}

func (m *Model) Init() tea.Cmd {
	return textarea.Blink
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	tiCmd := m.inputList.Update(msg)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.component.OnWindowSizeMsg(msg)

		m.jsonTa.SetWidth(msg.Width)
		m.jqTa.SetWidth(msg.Width)

	case tea.KeyMsg:
		return m.component.OnKeyPress(msg, m, m.parent, m.inputList)
	}

	return m, tiCmd
}

func (m *Model) View() string {
	s := style.Section("JSON:") + "\n"
	s += m.jsonTa.View() + "\n\n"
	s += style.Section("JQ Query:") + "\n"
	s += m.jqTa.View() + "\n\n"
	s += style.Section("Result:") + "\n"

	if m.err != nil {
		s += style.Error(m.err.Error())
	} else {
		if m.query == nil {
			b, err := json.MarshalIndent(m.json, "", "  ")
			if err != nil {
				m.err = err
				return m.View()
			}
			s += string(b)
		} else {
			var j any
			err := json.Unmarshal(m.json, &j)
			if err != nil {
				m.err = err
				return m.View()
			}

			iter := m.query.Run(j)
			for {
				v, ok := iter.Next()
				if !ok {
					break
				}
				if err, ok := v.(error); ok {
					m.err = err
					return m.View()
				}

				b, err := json.MarshalIndent(v, "", "  ")
				if err != nil {
					m.err = err
					return m.View()
				}
				s += string(b) + "\n"
			}
		}
	}

	return m.component.RenderView(
		"JSON Explorer", s, [][]key.Binding{},
	)
}

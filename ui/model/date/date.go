package date

import (
	"fmt"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dbarrosop/ttoys/ui/component"
	"github.com/dbarrosop/ttoys/ui/style"
)

type Model struct {
	component *component.Component
	parent    tea.Model

	inputList  *component.InputList
	rfc1123zTa textarea.Model
	rfc822z    textarea.Model
	rfc3339    textarea.Model
	unixTa     textarea.Model

	time time.Time
	err  error
}

func getTextArea() textarea.Model {
	ta := textarea.New()
	ta.Prompt = ""
	ta.ShowLineNumbers = false
	ta.SetHeight(1)
	ta.SetWidth(40)
	ta.KeyMap.InsertNewline = key.NewBinding(key.WithDisabled())
	ta.CharLimit = 31

	return ta
}

func New(comp *component.Component, parent tea.Model) *Model {
	rfc1123z := getTextArea()
	rfc1123z.Focus()
	rfc822z := getTextArea()
	rfc3339 := getTextArea()
	unixTa := getTextArea()

	m := &Model{
		component:  comp,
		parent:     parent,
		unixTa:     unixTa,
		rfc1123zTa: rfc1123z,
		rfc822z:    rfc822z,
		rfc3339:    rfc3339,
		inputList:  component.NewInputList(),
		time:       time.Now(),
	}

	m.updateTime()

	m.inputList.Add(
		component.Input{
			Input: &m.rfc1123zTa,
			AcceptAction: func() {
				m.time, m.err = time.Parse(time.RFC1123Z, m.rfc1123zTa.Value())
				if m.err != nil {
					return
				}
				m.updateTime()
			},
		},
	)

	m.inputList.Add(
		component.Input{
			Input: &m.rfc822z,
			AcceptAction: func() {
				m.time, m.err = time.Parse(time.RFC822Z, m.rfc822z.Value())
				if m.err != nil {
					return
				}
				m.updateTime()
			},
		},
	)

	m.inputList.Add(
		component.Input{
			Input: &m.rfc3339,
			AcceptAction: func() {
				m.time, m.err = time.Parse(time.RFC3339, m.rfc3339.Value())
				if m.err != nil {
					return
				}
				m.updateTime()
			},
		},
	)

	m.inputList.Add(
		component.Input{
			Input: &m.unixTa,
			AcceptAction: func() {
				var d int64
				d, m.err = strconv.ParseInt(m.unixTa.Value(), 10, 64)
				if m.err != nil {
					return
				}
				t := time.UnixMilli(d)
				m.err = nil
				m.time = t
				m.updateTime()
			},
		},
	)

	return m
}

func (m *Model) updateTime() {
	m.rfc1123zTa.SetValue(m.time.Format(time.RFC1123Z))
	m.rfc822z.SetValue(m.time.Format(time.RFC822Z))
	m.rfc3339.SetValue(m.time.Format(time.RFC3339))
	m.unixTa.SetValue(fmt.Sprintf("%d", m.time.UnixMilli()))
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	tiCmd := m.inputList.Update(msg)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.component.OnWindowSizeMsg(msg)

	case tea.KeyMsg:
		return m.component.OnKeyPress(msg, m, m.parent, m.inputList)
	}

	return m, tiCmd
}

func (m *Model) View() string {
	bindings := [][]key.Binding{}

	var s string
	s += style.Section("RFC1123Z") + style.Help(" ("+time.RFC1123Z+")") + "\n"
	s += m.rfc1123zTa.View() + "\n"

	s += style.Section("RFC822Z") + style.Help(" ("+time.RFC822Z+")") + "\n"
	s += m.rfc822z.View() + "\n"

	s += style.Section("RFC3339") + style.Help(" ("+time.RFC3339+")") + "\n"
	s += m.rfc3339.View() + "\n"

	s += style.Section("Unix Time") + "\n"
	s += m.unixTa.View() + "\n"

	if m.err != nil {
		s += style.Error(m.err.Error())
	}

	return m.component.RenderView("Date Formatter", s, bindings)
}

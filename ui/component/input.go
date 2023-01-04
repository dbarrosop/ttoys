package component

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type Input struct {
	Input        *textarea.Model
	AcceptAction func()
}

type InputList struct {
	inputs []Input
	idx    int
}

func NewInputList() *InputList {
	return &InputList{
		inputs: make([]Input, 0),
	}
}

func (il *InputList) Add(input Input) {
	il.inputs = append(il.inputs, input)
}

func (il *InputList) Update(msg tea.Msg) tea.Cmd {
	if il == nil || len(il.inputs) == 0 {
		return nil
	}

	var tiCmd tea.Cmd
	*il.inputs[il.idx].Input, tiCmd = il.inputs[il.idx].Input.Update(msg)

	return tiCmd
}

func (il *InputList) Next() {
	if il == nil || len(il.inputs) == 0 {
		return
	}

	prev := il.inputs[il.idx]
	il.idx++
	if il.idx >= len(il.inputs) {
		il.idx = 0
	}
	cur := il.inputs[il.idx]

	cur.Input.Focus()
	prev.Input.Blur()
}

func (il *InputList) Prev() {
	if il == nil || len(il.inputs) == 0 {
		return
	}

	prev := il.inputs[il.idx]
	il.idx--
	if il.idx < 0 {
		il.idx = len(il.inputs) - 1
	}
	cur := il.inputs[il.idx]

	cur.Input.Focus()
	prev.Input.Blur()
}

func (il *InputList) Accept() {
	if il == nil || len(il.inputs) == 0 {
		return
	}

	il.inputs[il.idx].AcceptAction()
}

func (il *InputList) Reset() {
	if il == nil || len(il.inputs) == 0 {
		return
	}

	il.inputs[il.idx].Input.Reset()
}

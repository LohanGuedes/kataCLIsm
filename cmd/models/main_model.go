package models

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lohanguedes/kataclism/cmd/styles"
)

var content = `
  #Content
  `

type status int

const (
	kataList status = iota
	readme
)

type Model struct {
	list     list.Model
	err      error
	readme   viewport.Model
	focused  status
	loaded   bool
	quitting bool
}

func New() *Model {
	return &Model{}
}

func (m *Model) Next() {
	if m.focused == readme {
		m.focused = kataList
	} else {
		m.focused++
	}
}

func (m *Model) Prev() {
	if m.focused == kataList {
		m.focused = readme
	} else {
		m.focused--
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.InitKatas(msg.Width, msg.Height)
			m.InitViewport(msg.Width, msg.Height)
			m.loaded = true
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "left", "h":
			m.Prev()
		case "right", "l":
			m.Next()
		}
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	m.readme.SetContent(m.getCurrentItemReadme())
	return m, cmd
}

func (m Model) View() string {
	if m.quitting {
		return ""
	}
	if m.loaded {
		kataView := m.list.View()
		readmeView := m.readme.View()
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			styles.FocusedStyle.Render(kataView),
			styles.FocusedStyle.Render(readmeView),
		)
	} else {
		return "loading..."
	}
}

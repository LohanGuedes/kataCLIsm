package models

import (
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

// handle error in a better way
func (m *Model) getCurrentItemReadme() string {
	kata, ok := m.list.SelectedItem().(Kata)
	if !ok {
		panic("Error creating reader")
	}
	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(m.readme.Width),
	)
	if err != nil {
		panic(err)
	}

	f, err := os.ReadFile(kata.readmePath)
	if err != nil {
		panic(err)
	}

	str, err := renderer.Render(string(f))
	if err != nil {
		panic(err)
	}

	return str
}

func (m *Model) InitViewport(width, height int) {
	vp := viewport.New(width/4*2, height-4)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)

	vp.SetContent(m.getCurrentItemReadme())
	m.readme = vp
}

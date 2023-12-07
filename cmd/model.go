package main

import (
	"fmt"

	"github.com/fatih/color"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	err      error
	selected map[int]struct{}
	katas    []Kata
	cursor   int
}

func initialModel() model {
	return model{
		katas:    []Kata{{title: "Kata 1"}, {title: "Kata 2"}, {title: "Kata 3"}},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.katas) - 1
			}
		case "down", "j":
			if m.cursor < len(m.katas)-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Select your Kata:\n\n"

	for i, kata := range m.katas {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		if cursor == ">" && checked != "x" {
			s += color.GreenString("%s [%s] %s\n", cursor, checked, kata.title)
		} else if checked == "x" && cursor != ">" {
			s += color.HiBlueString("%s [%s] %s\n", cursor, checked, kata.title)
		} else {
			s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, kata.title)
		}
	}

	s += "\t\tPress q to quit.\n\n"
	return s
}

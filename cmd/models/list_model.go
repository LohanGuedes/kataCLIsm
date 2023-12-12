package models

import (
	"github.com/charmbracelet/bubbles/list"
)

type Kata struct {
	title       string
	description string
	readmePath  string
}

func (k Kata) FilterValue() string {
	return k.title
}

// Implements defaultDelegate
func (k Kata) Title() string {
	return k.title
}

func (k Kata) Description() string {
	return k.description
}

func (k Kata) GetReadmePath() string {
	return k.readmePath
}

func (m *Model) InitKatas(width, height int) {
	katas := list.New([]list.Item{}, list.NewDefaultDelegate(), width/4*2, height-4)
	katas.SetShowHelp(false)
	katas.InfiniteScrolling = true
	m.list = katas

	m.list.Title = "Available Katas:"
	m.list.SetItems([]list.Item{
		Kata{
			title:       "buy milk",
			description: "strawberry milk",
			readmePath:  `./src/kata1.md`,
		},
		Kata{
			title:       "fold laundry",
			description: "or wear wrinkly t-shirts",
			readmePath:  `./src/kata2.md`,
		},
		Kata{
			title:       "eat sushi",
			description: "negitoro roll, miso soup, rice",
			readmePath:  `./src/kata3.md`,
		},
	})
}

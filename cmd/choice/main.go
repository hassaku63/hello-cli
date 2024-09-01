package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/hassaku63/hello-cli/internal/cli"
)

type Choice struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Model struct {
	Choices     []Choice         `json:"choices"`
	Cursor      int              `json:"cursor"`
	Selected    map[int]struct{} `json:"selected"`
	Canceled    bool             `json:"canceled"`
	IsSubmitted bool             `json:"isSubmitted"`
}

var _ tea.Model = (*Model)(nil)

func initialModel() Model {
	return Model{
		// Our to-do list is a grocery list
		Choices: []Choice{
			{Id: 1, Name: "Buy carrots"},
			{Id: 2, Name: "Buy celery"},
			{Id: 3, Name: "Buy kohlrabi"},
		},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		Selected: make(map[int]struct{}),
	}
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyUp:
			if m.Cursor > 0 {
				m.Cursor--
			}
		case tea.KeyDown:
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}
		case tea.KeySpace:
			_, ok := m.Selected[m.Cursor]
			if ok {
				delete(m.Selected, m.Cursor)
			} else {
				m.Selected[m.Cursor] = struct{}{}
			}
		case tea.KeyEnter:
			m.IsSubmitted = true
			return m, tea.Quit
		}

		switch msg.String() {
		case "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "j":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}
		case "q":
			m.Canceled = true
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() string {
	s := "What should we buy at the market?\n\n"

	for i, choice := range m.Choices {
		cursor := " "
		if m.Cursor == i {
			cursor = cli.Bold(cli.Green(">"))
		}

		checked := " "
		checkedStr := cli.Blue(fmt.Sprintf("[%s]", checked))
		if _, ok := m.Selected[i]; ok {
			checked = "x"
			checkedStr = cli.Bold(cli.Blue(fmt.Sprintf("[%s]", checked)))
		}

		s += fmt.Sprintf("%s %s (id=%d) %s\n", cursor, checkedStr, choice.Id, choice.Name)
	}

	s += "\nPress Enter to submit.\nPress q to quit.\n"

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	model, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	m := model.(Model)

	selected := make([]Choice, 0, len(m.Selected))
	for c := range m.Choices {
		if _, ok := m.Selected[c]; ok {
			selected = append(selected, m.Choices[c])
		}
	}

	fmt.Println("Selected:")
	for _, choice := range selected {
		fmt.Printf("  %s: %s\n", cli.Blue(fmt.Sprintf("%d", choice.Id)), choice.Name)
	}
}

package main

import (
	"encoding/json"
	"fmt"

	"net/http"
	"net/url"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Item interface {
	Title() string       // shown in the list
	Description() string // optional, used by delegate
	FilterValue() string // used for filtering
}

type fruit struct {
	name string
}

func (f fruit) Title() string {
	return f.name
}

func (f fruit) Description() string {
	return "" // optional
}

func (f fruit) FilterValue() string {
	return f.name
}

type focus int

const (
	focusInput focus = iota
	focusList
)

type animeMsg struct {
	_id       string
	character string
	show      string
	quote     string
}

type errMsg error

type model struct {
	input   textinput.Model
	loading bool
	anime   *animeMsg
	err     error

	list          list.Model
	allItems      []list.Item
	filteredItems []list.Item

	focus focus
}

func fetchAnimeQuote(character string) tea.Cmd {
	return func() tea.Msg {

		escaped := url.QueryEscape(character)

		url := fmt.Sprintf("https://yurippe.vercel.app/api/quotes?character=%s&random=1", escaped)
		resp, err := http.Get(url)
		if err != nil {
			return errMsg(err)
		}
		if resp.StatusCode != http.StatusOK {
			return errMsg(fmt.Errorf("API returned %s", resp.Status))
		}

		defer resp.Body.Close()

		var data []struct {
			Quote     string `json:"quote"`
			Character string `json:"character"`
			Show      string `json:"show"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return errMsg(err)
		}

		return animeMsg{
			show:      data[0].Show,
			character: data[0].Character,
			quote:     data[0].Quote,
		}
	}
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter character name"
	ti.Focus()
	ti.CharLimit = 64
	ti.Width = 20

	items := []list.Item{
		fruit{name: "Apple"},
		fruit{name: "Banana"},
		fruit{name: "Cherry"},
		fruit{name: "Date"},
		fruit{name: "Elderberry"},
		fruit{name: "Fig"},
		fruit{name: "Grapes"},
	}

	l := list.New(items, list.NewDefaultDelegate(), 50, 10)
	l.Title = "Fruits"

	return model{
		input:         ti,
		list:          l,
		allItems:      items,
		filteredItems: items,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	if m.focus == focusInput {
		m.input, cmd = m.input.Update(msg)
	} else {
		m.list, cmd = m.list.Update(msg)
	}

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "enter":
			character := m.input.Value()
			m.loading = true
			m.anime = nil
			m.err = nil
			return m, fetchAnimeQuote(character)
		case "q":
			return m, tea.Quit

		case "tab":
			if m.focus == focusInput {
				m.focus = focusList
				m.input.Blur()
				m.list.SetShowStatusBar(true)
			} else {
				m.focus = focusInput
				m.input.Focus()
				m.list.SetShowStatusBar(false)
			}
		}
	case animeMsg:
		m.loading = false
		m.anime = &msg

	case errMsg:
		m.loading = false
		m.err = msg

	}

	return m, cmd
}

func (m model) View() string {

	if m.loading {
		return m.input.View() + "\n\nLoading ..."
	}

	if m.err != nil {
		return m.input.View() + "\n\nError: " + m.err.Error()
	}

	if m.anime != nil {
		return fmt.Sprintf(
			"%s\n\nShow: %s\nCharacter: %s\n Quote: %s\n%s\n\nPress q to quit",
			m.input.View(),
			m.anime.show,
			m.anime.character,
			m.anime.quote,
			m.list.View(),
		)
	}

	return fmt.Sprintf(
		"%s\n\\n%s\n\nPress q to quit",
		m.input.View(),

		m.list.View(),
	)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package menu

import (
	"fmt"
	"sort"
)

type MenuItem struct {
	Name   string
	Action func()
	Order  int
}

type Menu struct {
	Title      string
	Items      []MenuItem
	ExitOption int
}

const defaultExitOption = 0
const invalidOptionMessage = "Invalid option. Please enter one of the listed numbers."

func NewMenu(title string, exitOption *int) *Menu {
	exitValue := defaultExitOption
	if exitOption != nil {
		exitValue = *exitOption
	}
	return &Menu{
		Title:      title,
		ExitOption: exitValue,
	}
}

func (m *Menu) AddItem(name string, action func(), order int) error {
	if order == m.ExitOption {
		return fmt.Errorf("%s has the same number as the exit option", name)
	}
	for _, item := range m.Items {
		if item.Order == order {
			return fmt.Errorf("duplicate order number: %d", order)
		}
	}
	m.Items = append(m.Items, MenuItem{
		Name:   name,
		Action: action,
		Order:  order,
	})
	return nil
}

func (m *Menu) Show() {
	var selected int
	for {
		clearScreen()
		mappedOptions := m.printMenu()
		if _, err := fmt.Scanln(&selected); err != nil {
			fmt.Println(invalidOptionMessage)
			fmt.Scanln()
			continue
		}
		if selected == m.ExitOption {
			fmt.Println("Exiting...")
			return
		}
		if item, ok := mappedOptions[selected]; ok {
			item.Action()
		} else {
			fmt.Println(invalidOptionMessage)
			fmt.Scanln()
			continue
		}
	}
}

func (m *Menu) printMenu() map[int]MenuItem {
	sortedOptions := m.orderedOptions()
	mapped := make(map[int]MenuItem)
	fmt.Printf("=== %s ===\n", m.Title)
	for _, item := range sortedOptions {
		fmt.Printf("%d. %s\n", item.Order, item.Name)
		mapped[item.Order] = item
	}
	fmt.Printf("%d. Exit\n", m.ExitOption)
	fmt.Print("> ")
	return mapped
}

func (m *Menu) orderedOptions() []MenuItem {
	sorted := make([]MenuItem, len(m.Items))
	copy(sorted, m.Items)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Order < sorted[j].Order
	})
	return sorted
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

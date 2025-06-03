package menu

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
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
	scanner := bufio.NewScanner(os.Stdin)

	for {
		clearScreen()
		mappedOptions := m.printMenu()

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				log.Printf("Input error: %v", err)
			}
			return
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			fmt.Println(invalidOptionMessage)
			continue
		}

		var selected int
		if _, err := fmt.Sscanf(input, "%d", &selected); err != nil {
			fmt.Println(invalidOptionMessage)
			WaitForEnter()
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
			WaitForEnter()
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

func GetConfirmation(prompt string) bool {
	fmt.Print(prompt)

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return false
	}

	response := strings.ToLower(strings.TrimSpace(scanner.Text()))
	return response == "y" || response == "yes"
}

func WaitForEnter() {
	fmt.Println("Press Enter to continue...")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}

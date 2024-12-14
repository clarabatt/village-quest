package tutorial

import (
	"fmt"
	"strconv"
)

type Vertex struct {
	Lat, Long float64
}

type Title struct {
	Title        string
	NumCaracters int
}

type AbstractTitle interface {
	description()
	print()
	addSufix(sufix string)
}

func (t Title) description() {
	fmt.Println("Title: ", t.Title+" has "+strconv.Itoa(t.NumCaracters)+" caracters")
}

func (t Title) print() {
	fmt.Println("A print from the method: ", t.Title)
}

func (t *Title) addSufix(sufix string) {
	t.Title = t.Title + " " + sufix
	t.NumCaracters = len(t.Title)
	fmt.Println("New Title: ", t.Title)
}

func classes() {
	fmt.Println("\n* Classes *")
	name := "Village Quest"
	title := Title{name, len(name)}

	title.print()

	title.addSufix(" - The Game")
	fmt.Println("Title: ", title.Title)
	title.description()

	var n AbstractTitle
	n = &title

	n.print()

}

func Tutorial() {
	fmt.Println("\n* Tutorial *")
	z, x, c, v := variables()
	fmt.Println(z, x, c, v, " <- variables")

	loops()

	conditions(z, x)

	pointers(z)

	arrays()

	dictionaries()
}

func dictionaries() {
	fmt.Println("\n* Dictionaries *")
	m := make(map[int]string)
	m[1] = "One"
	m[2] = "Two"

	fmt.Println("Map: ", m)
	fmt.Println("Map[1]: ", m[1])

	m2 := map[string]Vertex{
		"Toronto":   {43.70011, -79.4163},
		"Vancouver": {49.2827, -123.1207},
		"Montreal":  {45.50884, -73.58781},
	}

	fmt.Println("Map2: ", m2["Montreal"])

}

func arrays() {
	fmt.Println("\n* Arrays *")
	var a [5]int
	fmt.Println("Empty: ", a)

	odds := [6]int{1, 3, 5, 7, 9, 11}
	even := []int{2, 4, 6, 8, 10}

	fmt.Println("Odds: ", odds[2:3])
	fmt.Println("Even: ", even)

	even = append(even, 12)

	fmt.Println("Even: ", even)

	for i, number := range even {
		fmt.Println("Index: ", i, "Number: ", number)
	}

	for _, number := range odds {
		fmt.Println("Number: ", number)
	}

}

func pointers(z int) {
	fmt.Println("\n* Pointers *")
	p := &z
	fmt.Println("Pointer: ", p)
	fmt.Println("Value P: ", *p)

	*p = 100
	fmt.Println("Value Z: ", z)
}

func variables() (int, int, int, int) {
	fmt.Println("\n* Variables *")
	var i, j int = 11, 22
	k, l := 33, 44

	str := "Hello, World!"

	fmt.Println(i, j)
	fmt.Printf("%v - %T\n", k, k)
	fmt.Printf("%v - %T\n", l, l)
	fmt.Printf("%v - %T\n", str, str)

	return i, j, k, l
}

func loops() {
	fmt.Println("\n* Loops *")

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i + 10
	}

	fmt.Println("Sum:  ", sum)

	for sum < 200 { // while loop
		sum += 10
	}

	fmt.Println("Sum:  ", sum)
}

func conditions(x, y int) {
	defer fmt.Println("End of conditions")
	fmt.Println("\n* Conditions *")

	if x > y {
		fmt.Println("x is greater than y")
	} else {
		fmt.Println("x is less than or equal to y")
	}

	if v := x + 10; v > y {
		fmt.Println("x + 10 is greater than y")
	}

}

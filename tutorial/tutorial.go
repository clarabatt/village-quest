package tutorial

import (
	"fmt"
)

func Tutorial() {
	fmt.Println("* Tutorial *")
	fmt.Println(variables())
	loops()
}

func variables() (int, int, int, int) {
	fmt.Println("* Variables *")
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
	fmt.Println("* Loops *")

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i + 10
	}

	fmt.Println("Sum:  ", sum)

	for sum < 200 {
		sum += 10
	}

	fmt.Println("Sum:  ", sum)
}

package tutorial

import (
	"fmt"
)

func Tutorial() {
	fmt.Println("\n* Tutorial *")
	z, x, c, v := variables()
	fmt.Println(z, x, c, v, " <- variables")
	loops()
	conditions(z, x)
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

package main

import (
	"fmt"
)

func tutorial() {
	var i, j int = 11, 22
	k, l := 33, 44

	fmt.Println(i, j)
	fmt.Printf("%v - %T\n", k, k)
	fmt.Printf("%v - %T\n", l, l)
}

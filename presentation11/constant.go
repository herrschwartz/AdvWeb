package main

import "fmt"

const (
	foos = iota
	roh  = iota
	dah  = iota
)

func main() {
	const age int = 23
	fmt.Println(age)

	var house int = 3422

	fmt.Println(&house)

	var input string
	fmt.Scanln(input)
	fmt.Println(input)

}

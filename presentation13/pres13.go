package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	i := []int{1,2,3,4}
	greatestV(2,3,4,5)
	greatestS(i)

	var s string
	fmt.Scanln(s)
	fmt.Println(s)
}

func greatestV(items ...int) {
	max := 0
	for _, item := range items {
		if item > max {
			max = item
		}
	}
	fmt.Print("Largest: ")
	fmt.Println(max)
}

func greatestS([]int items) {
	max := 0
	for _, item := range items {
		if item > max {
			max = item
		}
	}
	fmt.Print("Largest: ")
	fmt.Println(max)
}

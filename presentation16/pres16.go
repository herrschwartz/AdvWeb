package main

import "fmt"

func main() {
	evenOrOdd(0)
	fmt.Println(evenOrOdd(1))
	greatest(1, 2, 3, 5)
}

func evenOrOdd(x int) (int, bool) {
	div := x / 2
	var dobs float64 = float64(x) / 2
	if dobs > float64(div) {
		return div, false
	} else {
		return div, true
	}
}

func greatest(items ...int) {
	max := 0
	for _, item := range items {
		if item > max {
			max = item
		}
	}
	fmt.Print("Largest: ")
	fmt.Println(max)
}

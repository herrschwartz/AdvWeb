package main

import "fmt"

func main() {

	// var num1, num2 int = 0
	// fmt.Sscanf(&num1, "%d")
	// fmt.Sscanf(&num2, "%d")
	// num3 := num1 % num2
	// fmt.Println(num3)

	for i := 0; i < 1000; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz")
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
		}
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("FizzBuzz")
		}
	}

	sum := 0
	for i := 0; i < 1000; i++ {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}

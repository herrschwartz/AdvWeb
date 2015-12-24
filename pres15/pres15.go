package main

import "fmt"

type customer struct {
	fname string
	lname string
	phone string
	age   int
}

func main() {
	intslice := []int{1, 2, 3, 4, 5}
	fmt.Println(intslice)
	strslice := []string{"hi", "mom"}
	fmt.Println(strslice)
	blah := make([]int, 5, 10)
	blah = append(blah, 2)
	fmt.Println(blah)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("len:", len(n))
	delete(n, "bar")

	ptr := new(int)
	fmt.Println(ptr)
	fmt.Println(*ptr)

	ptr2 := new(string)
	fmt.Println(ptr2)
	fmt.Println(*ptr2)

	ptr3 := new(bool)
	fmt.Println(ptr3)
	fmt.Println(*ptr3)

	cust1 := customer{fname: "josh", lname: "adams", phone: "555-555-5555", age: 22}
	cust2 := customer{fname: "josh", lname: "parker", phone: "556-555-5555", age: 23}
	fmt.Println(cust1)
	fmt.Println(cust2)
	fmt.Println(cust1.fname)
	//no
	//no

}

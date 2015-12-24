package main

import "fmt"

func main() {
	fmt.Println("Hello \t World \n")
	s := "GO HAVE A CAKE"
	l := " PLZ"
	fmt.Println(len(s))
	fmt.Println(s[2])
	fmt.Println(s[:3])
	fmt.Println(s + l)

	i := 5
	p := float64(i)
	j := int(p)
	fmt.Println(i)
	fmt.Println(p)
	fmt.Println(j)

	b := []byte(s)
	c := string(b)
	fmt.Println(b)
	fmt.Println(c)
}

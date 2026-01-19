package main

import "fmt"

func main() {
	fmt.Println("Hello its me Aftab")
	fmt.Println("Some changes")
	a := 10
	b := 20

	c := a + b
	fmt.Println(c)
	fmt.Println(a + b)

	fmt.Println("Hello World")

	if a < b {
		fmt.Println("B is big")
	}

	if c > b {
		fmt.Println("C is big")
	}
}

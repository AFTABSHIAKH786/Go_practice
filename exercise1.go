package main

import "fmt"

func main() {
	fmt.Println("Hello its me Aftab")

	a := 10
	b := 20

	c := a + b
	print(c)
	print(a + b)

	print("Hello World")

	if a < b {
		print("B is big")
	}

	if c > b {
		print("C is big")
	}
}

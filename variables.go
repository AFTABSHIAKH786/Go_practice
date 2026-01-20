package main

import (
	"fmt"
)

func main() {
	age := 23

	fmt.Println("I am", age, "years old")

	// Signed integers can store both negative and positive numbers (and zero).
	var i int = 10
	fmt.Println(i)

	// Unsigned integers can only store non-negative numbers (and zero).
	var u uint = 20
	fmt.Println(u)

	// Floats can store fractional numbers.
	// Float 64 has high precision
	var f float64 = 10.5341531351313123
	fmt.Println(f)

	// Float 32 has low precision this will result in low decimals numbers
	var lf float32 = 10.312334144342424134231
	fmt.Println(lf)

	// Boolean values hold true and false
	var isActive bool = true
	var isNotActive bool = false
	fmt.Println(isActive, isNotActive)

	// Complex numbers (we will never use these)
	var CN128 complex128 = 3 + 4i
	var CN64 complex64 = 3 - 5i
	fmt.Println(CN64, CN128)

	//Short hand decleration :=
	var full int = 20
	var small = 20
	short := 20
	fmt.Println(full, small, short)

	//Multi line variable declaration
	var apple, oranges = 20, 30
	fmt.Println(apple, oranges)
	var apple1, oranges1 = "Apple\n", "Oranges\n"
	fmt.Println(apple1, oranges1)

	//Mixed type variable declaration
	var a, b, c = 10.10, 10, "10"
	fmt.Printf("type of a, b ,c are %T\n", a, b, c)
	fmt.Println("\n")

	//Constant variables
	const CNST = "Constant variable"
	fmt.Println(CNST)

	//Escape sequences
	fmt.Println("new line \n")
	fmt.Println("alert \a")
}

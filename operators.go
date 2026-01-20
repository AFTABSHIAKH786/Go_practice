package main

import "fmt"

func main() {
	fmt.Println("Operators")

	a, b := 10, 20
	//Arithmetic operators
	fmt.Println("A + B ", a+b)
	fmt.Println("A - B ", a-b)
	fmt.Println("A * B ", a*b)
	fmt.Println("A / B ", a/b)
	fmt.Println("A % B ", a%b)
	a++
	fmt.Println("A++ ", a)
	a--
	fmt.Println("A-- ", a)

	//Relational Operators
	fmt.Println("A == B ", a == b)
	fmt.Println("A != B ", a != b)
	fmt.Println("A > B ", a > b)
	fmt.Println("A < B ", a < b)

	//Logical Operators
	C, D := true, false
	fmt.Println("C && D ", C && D)  //Both conditions should be true
	fmt.Println("C || D ", C || D)  //One of the conditions should be true
	fmt.Println("!C && !D", !C, !D) //Not conditions is true then false vice versa

	//Assignment Operators
	E, F := 20, 30
	E += F
	fmt.Println(E)
	E -= F //here 50 - 20 will happen
	fmt.Println(E)
}

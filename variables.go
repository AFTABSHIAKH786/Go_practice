package main

import "fmt"

func main() {
	var name = "Hello"

	fmt.Println(name)

	age := 23

	fmt.Println("I am", age, "years old")

	var isTrue bool = false

	fmt.Println(isTrue)

	if isTrue != true {
		fmt.Println("Is True is false")
	} else {
		fmt.Println("Is True is True")
	}

}

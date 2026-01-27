package main

import (
	"fmt"
	"time"
)

func main() {

	month := "June"

	switch month {
	case "January":
		fmt.Println("Winter")
	case "August":
		fmt.Println("Rainy")
	case "June":
		fmt.Println("Sunny")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its weekend!")
	default:
		fmt.Println("Its working day --", time.Now().Weekday())
	}

	IntType := func(i interface{}) {
		// The type switch statement is used to select the type of the interface
		// value and execute different code depending on the type. The type
		// switch statement is similar to the regular switch statement, but
		// it is used to switch on the type of the value rather than the value
		// itself. The type switch statement is used to inspect the type of
		// an interface value and execute different code depending on the type.
		switch t := i.(type) {
		case int:
			fmt.Println("Its an interger", t)
		case string:
			fmt.Println("Its an string", t)
		default:
			fmt.Println("Other type", t)
		}
	}

	IntType("aftab")
	IntType(10)

}

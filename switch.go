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
}

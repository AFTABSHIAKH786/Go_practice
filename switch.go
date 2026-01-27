package main

import "fmt"

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
}

package main

import "fmt"

func main() {

	//if else condition
	age := 15

	if age >= 18 {
		fmt.Println("You are eligilble for vote")
	} else {
		fmt.Println("You are not eligible")
	}

	score := 9

	if score == 15 {
		fmt.Println("Grade A+")
	} else if score < 15 {
		if score <= 13 && score >= 11 {
			fmt.Println("Grade A")
		} else {
			fmt.Println("Grade C")
		}
	}
}

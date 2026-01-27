package main

import "fmt"

func main() {

	//for loops
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// //nested loops
	// for i := 1; i <= 10; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Println(i, j)
	// 	}
	// }

	//while loop
	// i := 1

	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	//for loop continue and break

	for i := 0; i <= 12; i++ {

		if i == 4 {
			continue
		}

		if i == 10 {
			break
		}
		fmt.Println(i)
	}

	//	//GOTO statements
	//
	// end:
	//
	//	fmt.Println("End")
	//	goto start
	//
	// start:
	//
	//	fmt.Println("start ")
	//	goto end
}

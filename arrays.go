package main

import "fmt"

func main() {

	//Interger array if set empty it will take 0 as default
	// var arr [10]int
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr)

	//For a boolen it will make it false by default
	var boolen [10]bool
	fmt.Println(boolen)

	//For a string based array it will take empty string as default
	var str [10]string
	fmt.Println(str)

	//Slices uninitilized arrays that is default set to nil (equivalent to null)
	//Slices are dynamic that can store dynamic amount to values init.
	var slic []int
	slic = append(slic, 1, 2, 3, 4, 5, 6, 9)
	fmt.Println(slic)
	fmt.Println(cap(slic))
}

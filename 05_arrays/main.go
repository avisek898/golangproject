package main

import "fmt"

func main() {
	//arrays
	var fruitArr [2]string
	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"
	fmt.Println(fruitArr)
	fmt.Println((fruitArr[1]))
	//Declare and assign arrays at the same time
	Array1 := [2]string{"aaa", "bb"}
	fmt.Println(Array1)

	//slices
	fruitslices := []string{"apple", "orange", "banana"}
	fmt.Println(fruitslices)
	//for length of a slice
	fmt.Println(len(fruitslices))

	//count the range
	fmt.Println(fruitslices[0:0])
}

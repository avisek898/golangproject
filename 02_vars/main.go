package main

import "fmt"

func main() {
	// using var
	var name string = "Abhishek"
	var age int64 = 37
	var ishor = true
	ishor = false
	// shorhand
	name1 := " andy"
	fmt.Println(name, name1, age, ishor)
	fmt.Printf("%T\n", ishor)
}

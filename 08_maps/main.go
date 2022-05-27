package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string)

	//  assign keyvalues
	emails["bob"] = "bob@gmail.co"
	emails["saile"] = "saile@yahoo.com"
	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println((emails["bob"]))

}

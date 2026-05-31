package main

import "fmt"

func main() {
	fmt.Println("study of conditional statements")

	loginCount := 25
	var result string

	if loginCount < 10 {
		result = "regular user"
	} else if loginCount > 10 {
		result = "watch out"
	} else {
		result = "exactly 10 login count"
	}

	fmt.Println(result)

}

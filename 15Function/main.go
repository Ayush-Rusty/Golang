package main

import "fmt"

func main() {
	fmt.Println("study of function concept")
	greeter()
	greeterTwo()

	result := adder(3, 5)

	fmt.Println("Result is:", result)

	proResult := proAdder(2, 5, 8, 4)
	fmt.Println("pro result is going to be: ", proResult)

}

func adder(valueOne int, valueTwo int) int {
	return valueOne + valueTwo
}

func proAdder(values ...int) int {
	total := 0

	for _, value := range values {
		total = total + value
	}
	return total
}

func greeter() {
	fmt.Println("Golang is the new era language")
}

func greeterTwo() {
	fmt.Println("another method")
}

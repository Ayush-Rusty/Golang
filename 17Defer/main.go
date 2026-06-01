package main

import "fmt"

func main() {
	fmt.Println("study of defer concept")
	defer fmt.Println("hello world")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

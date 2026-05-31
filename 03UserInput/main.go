package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to the world of strangers"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the value of the reader:")

	input, _ := reader.ReadString('\n')
	fmt.Println("something went wrong while publishing the code,", input)
}

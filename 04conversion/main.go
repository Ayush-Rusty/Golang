package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// STRINGS ARE BASICALLY USED FOR STRING MANIPULATION
func main() {
	fmt.Println("welcome to the pizza shop")
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("thanks for rating the pizza,", input)

	numberRating, error := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("added 1 to your rating:", numberRating+1)
	}
}

package main

import "fmt"

func main() {
	fmt.Println("study of loop concept")

	days := []string{"sunday", "tuesday", "wednesday", "friday", "saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for index := range days {
	// 	fmt.Printf(days[index])
	// }

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 2 {
			goto strange
		}

		if rougueValue == 5 {
			break
		}

		fmt.Println("value is:", rougueValue)
		rougueValue++
	}

strange:
	fmt.Println("jumping ata random place for my work")
}

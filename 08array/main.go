package main

import "fmt"

func main() {
	fmt.Println("welcome to the concept of array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "tomato"
	fruitList[2] = "watermelon"

	fmt.Println("fruits list is: ", fruitList)
	fmt.Println("fruit lsit is: ", len(fruitList))

	var vegList = [5]string{"potato", "rice", "chicken"}
	fmt.Println("veg list is :", vegList)
}

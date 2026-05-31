package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to the concept of slices")
	var fruitList = []string{"apple", "watermelon", "pineapple"}

	fruitList = append(fruitList, "Mango", "banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	//THIS LINE IS BASICALLY USED TO CREATE A SLICE
	highScores := make([]int, 4)

	highScores[0] = 200
	highScores[1] = 300
	highScores[2] = 400
	highScores[3] = 500

	highScores = append(highScores, 900, 600, 700)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	var courses = []string{"react", "javascript", "python", "golang", "ruby", "rust"}
	fmt.Println("list down all the courses,", courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}

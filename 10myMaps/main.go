package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)
	languages["JS"] = "javascript"
	languages["RB"] = "ruby"
	languages["PY"] = "python"

	fmt.Println("list of all the languages:", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("list of all the languages:", languages)

	// lOOPS
	for key, value := range languages {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}
}

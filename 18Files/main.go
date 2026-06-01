package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("study of file handling in Golang")
	content := "This needs to go in a file -where it belongs to"

	file, err := os.Create("./myfile.text")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("length is:", length)
	defer file.Close()
	ReadFile("./myfile.txt")
}

func ReadFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("text data is inside the file is \n", databyte)
}

package main

import "fmt"

func main() {
	fmt.Println("study of structs in golang")
	username := User{"ayush", "knightdevil45@gmal.com", true, 16}
	fmt.Println(username)
	fmt.Printf("username details are: %+v\n", username)
	fmt.Printf("Name is %v and email is %v.", username.Name, username.email)
}

type User struct {
	Name   string
	email  string
	status bool
	Age    int
}

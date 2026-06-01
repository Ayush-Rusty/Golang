package main

import "fmt"

func main() {
	fmt.Println("study of structs in golang")
	username := User{"ayush", "knightdevil45@gmal.com", true, 16}
	fmt.Println(username)
	fmt.Printf("username details are: %+v\n", username)
	fmt.Printf("Name is %v and email is %v.", username.Name, username.Email)
	username.GetStatus()
	username.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "knightdevil45@gmail.com"
	fmt.Println("email of this user is: ", u.Email)
}

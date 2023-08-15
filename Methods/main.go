package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMain() {
	u.Email = "test@go.dev"
	fmt.Println("main of user is ?: ", u.Email)
}

func main() {
	akash := User{"akash", "akash@gmail.com", true, 12}
	akash.GetStatus()
	akash.NewMain()
	fmt.Println(akash)
}

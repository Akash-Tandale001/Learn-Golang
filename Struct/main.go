package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {

	// no inheritance in golang no super or parent

	akash := User{"Akash", "akash@gmail.com", true, 12}
	fmt.Println(akash)
	fmt.Printf("akash more details are : %+v", akash)
	fmt.Println(akash.Email)

}

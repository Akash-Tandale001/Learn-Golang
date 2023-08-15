package main

import "fmt"

func main() {
	// var one *int
	// fmt.Println(one)

	myNumber := 12
	var ptr *int = &myNumber
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = *ptr *2;
	fmt.Println(myNumber)
}

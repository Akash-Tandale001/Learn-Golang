package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	loginCount := 223

	var result string

	if loginCount < 10 {
		result = "regular user"
	} else if loginCount > 10 && loginCount < 50 {
		result = "jo"
	} else {
		result = "something else"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}
	if num := 3; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is greater than 10")
	}

	// if err != nil{

	// }

	//switch case in go lang
	rand.Seed(time.Now().UnixNano())
	diceNum := rand.Intn(4) + 1

	switch diceNum {
	case 1:
		fmt.Println("number is 1")
	case 2:
		fmt.Println("number is 2")
	case 3:
		fmt.Println("number is 3")
		fallthrough
	case 4:
		fmt.Println("number is 4")
	default:
		fmt.Println("what was that")
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const LoginToken string = "dkfd" //pubilc

func main() {
	var username string = "name"

	fmt.Println(username)
	fmt.Printf("Variable is type : %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)

	var smallvalue int = 255
	fmt.Println(smallvalue)

	var smallFloat float64 = 255.23456456546546
	fmt.Println(smallFloat)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is type : %T \n", anotherVariable)

	// no var

	numberOfUser := 200000 //allow in only in function
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of pizza : ")
	input, err := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input, err)

	numRting, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Error in parsing the input")
	} else {
		fmt.Println("Added 1 to rating : ", numRting+1)
	}

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	createdDate := time.Date(2020, time.March, 10, 12, 21, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))

	

}

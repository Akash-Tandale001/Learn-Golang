package main

import "fmt"

func main() {
	defer fmt.Println("opps just checking")
	defer fmt.Println("world defer")
	fmt.Println("hello defer")
	mydefer()

}

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

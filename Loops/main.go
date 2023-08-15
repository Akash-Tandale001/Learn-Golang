package main

import "fmt"

func main() {

	days := []string{"sun", "mon", "tue", "wed", "fir", "sat"}

	fmt.Println(days)
	// 1 kind
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	//2 kind
	for i := range days {
		fmt.Println(days[i])
	}

	//3 kind
	for i, val := range days {
		fmt.Printf("index is %v and value is %v\n", i, val)
	}

	ho := 1
	for ho <= 10 {
		if ho==2{
			goto loc
		}
		if ho==5{
			break
		}
		fmt.Println(ho)
		ho++
	}

	// goto statement
	loc:
		fmt.Println("hello")
}

package main

import "fmt"

func main() {
	fmt.Println("main function")
	greeter()

	// this is wrong
	// func greeter1()  {
	// 	fmt.Println("sdf")
	// }
	// greeter1()

	result := adder(5, 6)
	fmt.Println("result : ", result)

	result1, mymessage, check := proadder(5, 6, 7, 8, 92, 3, 4, 5, 6)
	fmt.Println("result1 : ", result1)
	fmt.Println(mymessage)
	fmt.Println(check)




}

func adder(valone int, valtwo int) int {
	return valone + valtwo
}

func proadder(val ...int) (int, string, string) {
	total := 0
	for _, valu := range val {
		total += valu
	}
	return total, "hi proadder function", "ohh check checkking"
}

func greeter() {
	fmt.Println("Namestry from golang")
}

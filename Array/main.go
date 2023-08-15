package main

import "fmt"

func main() {
	//array in go
	var arr [3]string
	arr[0] = "Hello"
	arr[1] = "Akash"
	arr[2] = "How are you ?"

	fmt.Println(arr)
	fmt.Println(len(arr))

	var arr1 = [3]string{"ko", "go", "mo"}
	fmt.Println(arr1)

	//slice in go
	var arr2 = []string{"apple", "po", "io", "lop", "mo"}
	fmt.Println(arr2)
	fmt.Printf("%T", arr2)
	arr2 = append(arr2, "go","banana")
	fmt.Println(arr2)
	arr2 = append(arr2[:4])
	fmt.Println(arr2)

	//remove element from slice

	var arr3 = []string{"po","jo","lo","fo"};
	fmt.Println(arr3)
	var index int = 2;
	arr3=append(arr3[:index],arr3[index+1:]...);
	fmt.Println(arr3)
}

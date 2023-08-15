package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("file handing")
	content := "my out in file"
	file, err := os.Create("./myconfi.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is : ", length)
	defer file.Close()

	readFile("./myconfi.txt")

}

func readFile(filename string) {
	databite, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("text inside file : ", string(databite))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

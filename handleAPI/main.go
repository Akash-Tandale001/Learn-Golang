package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:1111/get"
	res, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println(res.StatusCode)
	fmt.Println(res.ContentLength)
	content, _ := ioutil.ReadAll(res.Body)

	var resString strings.Builder
	byteCount, _ := resString.Write(content)
	fmt.Println(byteCount)
	
	
	fmt.Println(resString.String())
	fmt.Println(string(content))
}

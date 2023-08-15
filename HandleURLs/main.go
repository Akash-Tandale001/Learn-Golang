package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?course=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparam := result.Query()
	fmt.Println(qparam)
	fmt.Println(qparam["course"])
	for _, val := range qparam {
		fmt.Println("Param is ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=Akash",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}

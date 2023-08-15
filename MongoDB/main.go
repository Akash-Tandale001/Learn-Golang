package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/Akash-Tandale001/MongodbWithGolang/route"
)

func main() {
	fmt.Println("Mongo")
	r := router.Router()
	fmt.Println("Server is getting started ....")	
	log.Fatal(http.ListenAndServe(":4000",r))
	fmt.Println("Listening at port 4000 ....")
}


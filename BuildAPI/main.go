package main

import (
	"math/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware / helper -file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""

}

func main() {
	fmt.Println("API Creation")
	// create a new router instance with default settings
	r := mux.NewRouter()

	//seeding of data
	courses = append(courses,
		Course{
			CourseId:    "1",
			CourseName:  "reactjs",
			CoursePrice: 299,
			Author:      &Author{"Akash", "akash@gmail.com"}})
	courses = append(courses,
		Course{
			CourseId:    "2",
			CourseName:  "Embar",
			CoursePrice: 499,
			Author:      &Author{"Crowdstrike", "crowdstrike@gmail.com"}})

	//routing
	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/course",addOneCourse).Methods("POST")
	r.HandleFunc("/updateCourse/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}",deleteOneCourse).Methods("DELETE")





	// listen to  a port
	log.Fatal(http.ListenAndServe(":4000", r)) 


}

//Controllers - file

// serve home route
func serveHome(writer http.ResponseWriter, reader *http.Request) {
	writer.Write([]byte("<h1>Welcome to my API</h1>"))
}

func getAllCourses(writer http.ResponseWriter, reader *http.Request) {
	fmt.Println("Get all Course")
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(courses)
}

func getOneCourse(writer http.ResponseWriter, reader *http.Request) {
	fmt.Println("Get one course")
	writer.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(reader)
	//loop through courses find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(writer).Encode(course)
			return
		}
	}
	json.NewEncoder(writer).Encode("No course found with given  id")
	return
}

func addOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("request Body cannot be null ")
		return
	}
	// what about -{}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Empty object passed in as a new course")
		return
	}

	//generate unique id, convert in string
	//append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("Content-Type", "application/json")

	//first grab id from req
	params := mux.Vars(r)
	// loop id remove add opration with myID from params

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// send a response when id is not found
	json.NewEncoder(w).Encode("course not found")
	return
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one course")
	w.Header().Set("Content-Type", "application/json")
	//grab the ID of the item to be deleted
	params := mux.Vars(r)
	for i, course := range courses {
		if courses[i].CourseId == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found")
	return
}

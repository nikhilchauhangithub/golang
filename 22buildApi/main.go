package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {
fmt.Println("Api in Golang")
r :=mux.NewRouter()

//seeding
courses =append(courses, Course{CourseId: "17",CourseName: "Golang", CoursePrice: 1000, Author: &Author{Fullname: "Nikhil Chauhan",Website: "github.com"}})

//routing

r.HandleFunc("/",serveHome).Methods("GET")
r.HandleFunc("/courses",getAllCourses).Methods("GET")
r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
r.HandleFunc("/course",createOneCourse).Methods("POST")
r.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT")
r.HandleFunc("/courses/{id}",deleteOneCourse).Methods("DELETE")

//listen to a port
log.Fatal(http.ListenAndServe(":4000",r))
}

//Model for course-file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake database
var courses []Course

func (c *Course) IsEmpty() bool {
	//middleware
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

//controller file
//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
w.Write([]byte("<h1>I am learning API in golang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses) // json.NewEncoder(w) creates a new json.Encoder instance that writes JSON data to the http.ResponseWriter instance w. By using an Encoder to encode the JSON data, the function can stream the JSON data directly to the response writer without buffering it in memory. This is more memory-efficient than buffering the JSON data in memory first and then writing it to the response writer, especially if the JSON data is large.
}

func getOneCourse(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("content-Type", "application/json")

	//grab id from request
	params :=mux.Vars(r) //it is used extract course id from user's requested url and send it to params

	//loop through courses, find matching id and return the response to user

	for _,course:=range courses{
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode(fmt.Sprintf("No such course found with given id:%v\n",params["id"]))
}

func createOneCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("create one course")
	w.Header().Set("Content-type", "application/json")

	//what if body is empty
	if r.Body==nil {
		json.NewEncoder(w).Encode("Empty field is not accepted")
	}

	//what if we got {}

	var course Course //This code block appears to be decoding JSON data from an HTTP request body into a Course struct, and then checking if the decoded Course instance is empty. If the Course is empty, the code block sends a response to the HTTP client indicating that no data was sent.
	json.NewDecoder(r.Body).Decode(&course)//By passing &course to the Decode() method, the method can decode the JSON data and store the resulting Course struct instance directly into the original course variable, which can then be used in subsequent code.
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("send some data")
		return
	}
	//generate unique id, then convert it to string

	rand.Seed(time.Now().UnixNano())
	course.CourseId=strconv.Itoa(rand.Intn(100))
	courses =append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	//first- grab id from req
	params:=mux.Vars(r)

	//loop, Id, remove, add with my ID

	for index,course:=range courses{
		if course.CourseId==params["id"] {
			courses =append(courses[:index], courses[index+1:]...)
var course Course
_=json.NewDecoder(r.Body).Decode(&course)
course.CourseId=params["id"]
courses=append(courses, course)
json.NewEncoder(w).Encode(course)
return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")
	params :=mux.Vars(r)

	//loop, id, remove id (:index, index+1)
	 
	for index, course := range courses{
		if course.CourseId==params["id"] {
			courses =append(courses[:index], courses[index+1:]...)
		}
		break
		}
		json.NewEncoder(w).Encode("course deleted successfully")
}
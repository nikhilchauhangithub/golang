package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

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
	return c.CourseId == "" && c.CourseName == ""
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
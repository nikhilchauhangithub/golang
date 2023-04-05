package main

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
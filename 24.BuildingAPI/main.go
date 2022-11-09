package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//creating a model for course  -file

type Course struct {
	CourseName string  `json:"coursename"`
	CourseId   string  `json:"courseid"`
	Price      int     `json:"price"`
	Author     *Author `json:"auhtor"`
}

type Author struct {
	Fullname string `json:"fullName"`
	Website  string `json:"website"`
}

//fake Db

var coursesDb []Course // having a slice of courses

// helper function or middleware function as a validation

func (c *Course) isEmpty() bool {

	if c.CourseId == "" || c.CourseName == "" {
		return false
	} else {
		return true
	}

}

func main() {

	fmt.Println(" welcome to the main function ")

	r := mux.NewRouter()

	//seeding of db in mian func
	coursesDb = append(coursesDb, Course{CourseId: "2", CourseName: "MyGoLang", Price: 23, Author: &Author{Fullname: "chandanGupta", Website: "chandanGupta.com"}})
	coursesDb = append(coursesDb, Course{CourseId: "2", CourseName: "MyGoLang", Price: 23, Author: &Author{Fullname: "chandanGupta", Website: "chandanGupta.com"}})

	//routing

	r.HandleFunc("/", serveHome).Methods("Get")
	r.HandleFunc("/getAllCourse", getAllCourses).Methods("Get")
	r.HandleFunc("/getOneCourse/{id}", getOneCourse).Methods("Get")

	r.HandleFunc("/creatCourse", createOneCourse).Methods("Post")

	// listening to the port
	log.Fatal(http.ListenAndServe(":3030", r))
}

//controllers

//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1> welcome to the home route </h1>"))

}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("welcome to the all courses api controller")

	// writing an header
	w.Header().Set("contentType", "application/json")

	// sending an json file

	json.NewEncoder(w).Encode(coursesDb)
	// this is like hey json encoder work with the write or w
	// working with encode()
	// whatever  inside the encode will be treate as json value
	//and sent back whoever is making request
}

// grabig only one course
func getOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println(" welcome to the get on course handler ")

	// grab id from the request

	params := mux.Vars(r)
	// this can also be done using url package
	// but mux is also providing the facility

	// find the matching id in arrayor db  and send it in response

	for _, element := range coursesDb {

		if element.CourseId == params["id"] {
			json.NewEncoder(w).Encode(element)
			return
		}

	}

	// nothing found out of the loop
	json.NewEncoder(w).Encode("no course Found ")

}

//create one course by having json data post Request

func createOneCourse(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("contentType", "application/json")

	// we wll be getting dat in json format

	// what if no body is there

	// we can have access to the body via r.body

	if r.Body == nil {
		json.NewEncoder(w).Encode(" please provide a body ")
	}

	// if {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	// if course.isEmpty() {
	// 	json.NewEncoder(w).Encode("no data is in bdy ")

	// 	return

	// }

	// courses variable will be having response of json   on decode(&course)

	// now have this course  create a unique id and push it into courses slice

	coursesDb = append(coursesDb, course)

	json.NewEncoder(w).Encode(coursesDb)

}

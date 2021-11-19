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

//Modal for course - file


type Course struct {

	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"courseprice"`
	Author *Author `json:"author"`
}


type Author struct {
	FullName string `json:"fullname"`
	Website string `json:"website"`
	
}

//make DB

var courses []Course

//middleware,helper - file

func (c *Course) IsEmpty() bool {
	//return c.CourseId == "" && c.CourseName == ""
	return   c.CourseName == ""
}

func main() {
	fmt.Println("API - Course JAY")
	r := mux.NewRouter()

	//sedding
	courses = append(courses,Course{CourseId:"4",CourseName: "NextJs",CoursePrice:45, Author : &Author{FullName : "JayChauhan",Website : "http://jay.com"}})
	courses = append(courses,Course{CourseId:"8",CourseName: "ReactJs",CoursePrice:450, Author : &Author{FullName : "JayChauhan",Website : "http://jay00.com"}})
 
	//routing
	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllCorses).Methods("GET")
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/course",createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}",deleteOneCourse).Methods("DELETE")

	//listen to a port


	log.Fatal(http.ListenAndServe(":8080", r))


}

//controller - file 

// serve home routes


func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by jat</h1>"))
}

func getAllCorses(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get All Course")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
fmt.Println("Get One Course")
w.Header().Set("Content-Type", "application/json")
// grab id from request

   param := mux.Vars(r)

   //loop through courses , find matching courses and return
   
   for _, c := range courses {
	   if c.CourseId == param["id"] {

		json.NewEncoder(w).Encode(c)

		return
		   
	   }
   }
   json.NewEncoder(w).Encode("No course found with given id")

   return
}

func createOneCourse(w http.ResponseWriter, r *http.Request){
	 fmt.Println("Create One Course")
      w.Header().Set("Content-Type", "application/json")
    // what if : body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data to the server")

	}
	// what about -- {}
     var course Course

	_= json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}
	// generate unique id, string
	//append to course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request){
 
   fmt.Println("update One Course")
      w.Header().Set("Content-Type", "application/json")

	  //first - grabe id from req
	   params := mux.Vars(r)


	   //loop,id.remove , add with ID my

	   for index,course := range courses {
		   if course.CourseId == params["id"] {
			   courses = append(courses[:index], courses[index+1:]...)
            var course Course
			   json.NewDecoder(r.Body).Decode(&course)
			   course.CourseId = params["id"]
			   courses =append(courses, course)
			   json.NewEncoder(w).Encode(course)
			   return
		   }
	   }

}


func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	 fmt.Println("delete One Course")
      w.Header().Set("Content-Type", "application/json")


	   params := mux.Vars(r)
     
	   //loop,id.remove , (index,index +1)

	    for index,course := range courses {
		   if course.CourseId == params["id"] {
			   courses = append(courses[:index], courses[index+1:]...)
			      json.NewEncoder(w).Encode("new list is ") 
			   json.NewEncoder(w).Encode(courses)              
			   break
		   }
		   
			   json.NewEncoder(w).Encode("not a walid ID")
		   
	   }

}
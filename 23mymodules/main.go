package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("heele mod in golang")
	greeter()
	 r := mux.NewRouter()
    r.HandleFunc("/",serveHome).Methods("GET")


     log.Fatal(http.ListenAndServe(":4000",r))
	//http.ListenAndServe(":4000",r)
}


func greeter(){
	fmt.Println("hey mod user")
}

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1> welcome to the go lang</h1>"))
}
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jay75chauhan/mongoapi/router"
)

    func main() {
	fmt.Println("hii mongo movie db")
    r :=router.Router()
	fmt.Println("Server is getting started...")
    log.Fatal(http.ListenAndServe(":4000",r))   
   fmt.Println("Server is getting started at 4000 now...")


}
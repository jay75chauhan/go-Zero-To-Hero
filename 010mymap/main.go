package main

import "fmt"

func main() {

	fmt.Println("hii map")


     la := make(map[string]string)

	 la["js"]= "javascript"

	 la["kt"]= "kotlin"

	 la["py"]= "pythone"

 
       fmt.Println("hii map list is",la)
       fmt.Println("hii map list is",la["kt"])


    delete(la, "kt")

 fmt.Println("hii map list is deleted",la)

//lop in goto

	for value,key := range la {

       fmt.Printf(" key: %v, value: %v\n", key, value)

	}





}
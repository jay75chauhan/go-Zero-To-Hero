package main

import "fmt"

func main() {

	fmt.Println("Starting")

	// var ptr *int 

	// fmt.Println("Starting", ptr)


	mynum := 2020

	var ptr = &mynum

	fmt.Println("Starting", ptr)
	fmt.Println("Starting", *ptr)

	*ptr = *ptr + 2 
	fmt.Println("Starting new", mynum)
	fmt.Println("Starting", *ptr)

}
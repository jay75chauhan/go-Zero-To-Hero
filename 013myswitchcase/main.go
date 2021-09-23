package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("hii switch case")

fmt.Println("rool the dice")
	rand.Seed(time.Now().UnixNano())
	diceNum := rand.Intn(6) + 1

	fmt.Println(  "the nuber is",  diceNum )


	switch diceNum {
		case 1: 
		fmt.Println("you can move : 1")
		case 2: 
		fmt.Println("you can move : 2")
		case 3: 
		fmt.Println("you can move : 3")
		case 4: 
		fmt.Println("you can move : 4")
		case 5: 
		fmt.Println("you can move : 5")
		case 6: 
	              	fmt.Println("you can move : 6")
		fallthrough
	    default:
     	fmt.Println("what a man this is wrong")

	}
}
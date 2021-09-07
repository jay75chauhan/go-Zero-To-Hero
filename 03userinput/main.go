package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    wel := "welcome to the go world"
	fmt.Println(wel)


	reader := bufio.NewReader(os.Stdin)
fmt.Println("enter the rating")


// comma ok // err err

input, _ := reader.ReadString('\n')
fmt.Println("Thanks for ", input)
fmt.Printf("Thanks for %T ", input)


}
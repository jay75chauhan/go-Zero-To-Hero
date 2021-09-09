package main

import (
	"fmt"
	"math/big"
	//"math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("hiiiiiiii")


	// var mynum1 int = 2

	// var mynum2 float64 = 2


	// fmt.Println("hiiiiii",mynum1 + int(mynum2))


	//random number generator


    //  rand.Seed(time.Now().UnixNano())

	// fmt.Println(rand.Intn(5) + 1)



	mynum,_ := rand.Int(rand.Reader,big.NewInt(5))
	fmt.Println(mynum)
}
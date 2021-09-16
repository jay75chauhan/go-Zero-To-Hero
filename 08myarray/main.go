package main

import "fmt"

func main() {

	fmt.Println("hiii array")



	var fru [4]string

	fru[0] = "apple"
	fru[1] = "apple 2"
	fru[3] = "apple 3"

	fmt.Println("apple is",fru)
	fmt.Println("apple is",len(fru))


var veg =  [4]string{"aple", "a","banana"}

 fmt.Println("veg is",veg)
 fmt.Println("veg is",len(veg))
}
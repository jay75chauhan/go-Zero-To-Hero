package main

import (
	"fmt"
	"sort"
)

func main() {


	
	fmt.Println("hii into slice")
    
	var frut = []string{"a", "b", "c", "d", "e"}

    fmt.Printf("hii into slice is %T \n ",frut)

       frut = append( frut,"rr")
       fmt.Println(frut)

	   frut = append(frut[1:2])
       fmt.Println(frut)



	   score := make([]int,4)

	   score[0] = 1
	   score[1] = 3
	   score[2] = 5
	   score[3] = 2
	//    score[4] = 255
	   
	score = append(score,44,66)

     fmt.Println(score)

    sort.Ints(score)
    fmt.Println(score)
    fmt.Println(sort.IntsAreSorted(score))

	//remove value from slice

	var course = []string{"r","h","y","r"}

    fmt.Println(course)

        var i int = 2
	course = append(course[:i], course[i+1:]...)




	fmt.Println(course)


	

}
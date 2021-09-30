package main

import "fmt"

func main() {
	fmt.Println("hii loop")

	// days := []string{"Sun", "Mon", "Tue", "Wed", "Thu"}


	// fmt.Println(days)
    // for i:=0; i<len(days); i++ {
    //   	fmt.Println(days[i])

	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }


	// for i, v := range days {
	// 	fmt.Printf("index is %v -- valie is %v",i,v)
	// }


	
	// for _, v := range days {
	// 	fmt.Printf("index is -- valie is %v",v)
	// }



	ro := 1
	for ro< 10 {
		   if ro == 3 {
			  goto ll
		  }
		if ro == 5 {
			// break
          
             ro++
			 continue
		}
	fmt.Println("value", ro)
	ro++
	}

	ll:
	fmt.Println("hii jay")
	  

}
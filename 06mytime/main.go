package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome my frend")

	present := time.Now()

	fmt.Println(present)


	fmt.Println(present.Format("01-02-2006 15:04:05 Monday"))

	creatT := time.Date(2020,time.December, 10,23,23,0,0, time.UTC)


	fmt.Println(creatT)
	fmt.Println(creatT.Format("01-02-2006 15:04:05 Monday"))

}

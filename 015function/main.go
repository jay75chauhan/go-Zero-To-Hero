package main

import "fmt"

func main() {

	fmt.Println("hiii")
      gg()

      result:=adder(2,4)
	  fmt.Println("hiii is result",result)



	  r3,m := proadder(2,5,5,5)
	  
	  fmt.Println("hiii is result",r3,m)
}

func proadder(values...int)(int,string){

	total:= 0

	for _,val := range values{
		total+= val
	}

	return total ,"hu jay"

}

func adder(v1 int, v2 int) int {
    return v1 + v2
}

func gg(){
	fmt.Println("hiii 5")
}
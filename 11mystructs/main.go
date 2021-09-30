package main

import "fmt"

func main() {
	fmt.Println("welcome to structs")

 jay := User{"JAY",44,"JAY@GMAIL.COM"}

 fmt.Println(jay)
 fmt.Printf("hiii : %+v\n",jay)
 fmt.Printf("name is :%v age is : %v ",jay.Name,jay.Age)
  

}

type User struct{ 
Name string
Age int
Email string
}
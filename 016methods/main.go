package main

import "fmt"

func main() {
	fmt.Println("welcome to structs")

 jay := User{"JAY",44,"JAY@GMAIL.COM"}

 fmt.Println(jay)
 fmt.Printf("hiii : %+v\n",jay)
//  fmt.Printf("name is :%v age is : %v ",jay.Name,jay.Age)
  
 jay.GetStatus()
 jay.NewMail()
 fmt.Printf("name is :%v age is : %v ",jay.Name,jay.Email)

}

type User struct{ 
Name string
Age int
Email string

}

func (u User) GetStatus(){

	fmt.Println("is user active",u.Email)

}

func (u User) NewMail(){

	
	u.Email = "r@gmain"


	fmt.Println("is user new email is",u.Email)

}
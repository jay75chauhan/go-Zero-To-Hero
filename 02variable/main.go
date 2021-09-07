package main

import "fmt"


const Login string = "Login"

func main() {
	var user string = "jay"
	fmt.Println(user)
	fmt.Printf("is type %T \n",user)

	var isuser bool = true
	fmt.Println(isuser)
	fmt.Printf("is type %T \n",isuser)

	var smallValue uint8 = 25
	fmt.Println(smallValue)
	fmt.Printf("is type %T \n",smallValue)

	var smallFloat float32 = 255.5555555555555
	fmt.Println(smallFloat)
	fmt.Printf("is type %T \n",smallFloat)


	// default value and some aliases

	var anotherValue int
	fmt.Println(anotherValue)
	fmt.Printf("is type %T \n",anotherValue )

	//implicit type

	var jay = "jay.tercjh"
	fmt.Println(jay)


// no var style

nuberOf := 7099

fmt.Println(nuberOf)


fmt.Println(Login )
fmt.Printf("is type %T \n",Login )


}
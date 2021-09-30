package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://lco.dev"

func main() {

	fmt.Println("hii web requ")

     rp,err  :=	http.Get(url)

	 if err != nil {
		 panic(err)
	 }

	 fmt.Printf("respons is %T\n",rp)

	defer rp.Body.Close()  // caller's responsibility to clone the conecttion


      databyte,err :=	ioutil.ReadAll(rp.Body) //

	   if err != nil {
		 panic(err)
	 }

	 fmt.Println("data is ",string(databyte))

}
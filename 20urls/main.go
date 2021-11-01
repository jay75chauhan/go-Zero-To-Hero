package main

import (
	"fmt"
	"net/url"
)


const  myurl  string = "http://lco.dev:3000/learn?corsename=reacjs&paymentid=rfjijf"

func main() {

	fmt.Println("hii url handling")
	fmt.Println(myurl)

	//parsing

	result,_ := url.Parse(myurl)


	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("type of Query is %T \n", qparams)


	fmt.Println(qparams["corsename"])


	for _,val := range qparams{
		fmt.Println("param is:",val)
	}

	partofUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/sucess",
		RawPath: "user=jay",
	}


	anotherURL := partofUrl.String()
	fmt.Println(anotherURL)
}
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("welcome to the web")
	GetRequest()

}


func GetRequest() {

	const url = "http://localhost:8000/get"

	respons,err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer respons.Body.Close()

	fmt.Println("status code",respons.StatusCode)
	fmt.Println("content leght",respons.ContentLength)

	 var responseString strings.Builder

	content,_ := ioutil.ReadAll(respons.Body)
	byteCod,_ := responseString.Write(content)


	fmt.Println("byteCount is",byteCod)
	fmt.Println(responseString.String())


	//fmt.Println(string(content))
}
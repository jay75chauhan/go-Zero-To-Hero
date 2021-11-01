package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("welcome to the web")
	//GetRequest()
	//PostRequest()
	PostForm()

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

func PostRequest() {
	const myurl = "http://localhost:8000/post"
    

    //face json payload

	requestBody := strings.NewReader(`
	{

		"name":"jay",
		"age":20,
	}
		`)

	respons,err :=	http.Post(myurl,"application/json" ,requestBody)

	if err != nil{
		panic(err)
	}

   defer respons.Body.Close()

   content,_ := ioutil.ReadAll(respons.Body)

   fmt.Println(string(content))
 



}

func PostForm(){

	const myurl = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}


	data.Add("first_name","jay")
	data.Add("last_name","chauhan")
       

	respons,err := http.PostForm(myurl,data)


	if err != nil {
		panic(err)
	}
	defer respons.Body.Close()
     
	content,_ := ioutil.ReadAll(respons.Body)

	fmt.Println(string(content))
}
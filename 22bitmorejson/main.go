package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"name"`
	Price    int
	Password string `json:"-"`
	Tags   []string `json:"tags,omitempty"`
}

func main() {

	fmt.Println("welcome to JSON ")


	EncodeJson()
}


func EncodeJson() {



	CoursesName := []course{
{"jay",66,"1234",[]string{"wev-dev","33"}},
{"jay",66,"1234",[]string{"wev-dev","33"}},
{"jay",66,"1234",nil},
}


  //package this data into JSON data

 finalJson,err := json.MarshalIndent(CoursesName, "", "\t ")


 if err != nil {
	 panic(err)
 }
 
 fmt.Printf("%s\n",finalJson)

}
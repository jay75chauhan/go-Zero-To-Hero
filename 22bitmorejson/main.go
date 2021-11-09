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


	//EncodeJson()
	DecodeJson()
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
 fmt.Println(finalJson)
 
 fmt.Printf("%s\n",finalJson)

}


func DecodeJson() {

	jsonDatafrom := []byte(`
	 {
	 	 "name": "jay",
	 	 "Price": 66,
	 	 "tags": [ "wev-dev","33" ]
	 }	
	`)

	var coursejay course 
	
	 checkValid :=  json.Valid(jsonDatafrom)



	 if checkValid {
             fmt.Println("JSON is walid")
			 json.Unmarshal(jsonDatafrom, &coursejay)
			 fmt.Printf("%#v\n",coursejay)
	 } else {
		  fmt.Println("JSON is NOT walid")
	 }



	 // some case where you just to add data to key values


	 var mydata map[string]interface{}
	 json.Unmarshal(jsonDatafrom, &mydata)
	 fmt.Printf("%#v\n",mydata)

	 for k, v := range mydata {
		 fmt.Printf("key : %v, value : %v and type : %T\n", k, v,v)
	 }

} 
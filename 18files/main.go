package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	 fmt.Println("")

	 content :="this is my new go file..."

     file,err :=	os.Create("./myfile.txt")


//   if err != nil {
// 	  panic(err)
//   }
         checkNilErr(err)

      
		 le,err:=    io.WriteString(file,content)
   
		 checkNilErr(err)
  
		 fmt.Println("lenth is",le)
   
		 defer file.Close()
   
		 readFill("./myfile.txt")
}





func readFill(filename string){
	
	databyte, err := ioutil.ReadFile(filename)
	
  
	checkNilErr(err)

	fmt.Println("data is inside file \n",string(databyte))

}


func checkNilErr(err error)  {

 
	if err != nil {

		panic(err)
  }

}
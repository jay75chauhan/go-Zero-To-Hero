package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"tests"}

var wg sync.WaitGroup //pointer

var mutex sync.Mutex //pointer

func main() {

//  go	greeter("hello")
// 	greeter("world")
         
       websitelist := []string{
		   "https://localhost.com",
		   "https://jay.com",
		   "https://fb.com",
		   "https://google.com",
		   "https://github.com",
		   "https://jaychauhan.tech",		
		}
 

		  for _, s := range websitelist {
			go  getStatusCode(s)
			wg.Add(1)
		  }

		  wg.Wait()
		  fmt.Println(signals)

}

// func greeter(s string) {

//  for i:=0;i<6;i++ {

// 	time.Sleep(1 * time.Second)
//      fmt.Println(s)
//  }	
// }


func getStatusCode(endpoint string)  {

	defer wg.Done()

  res,err:=	http.Get(endpoint)

  if err != nil {
	  fmt.Println("OOPS in endpoint")
  } else{        

	mutex.Lock()
	signals = append(signals, endpoint)
	mutex.Unlock()
	
	
	fmt.Printf("%d status code for %s\n",res.StatusCode,endpoint)
  }
}
package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Channels in go lange")
	myCh := make(chan int,2)

	wg:= &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println(<-myCh)

	wg.Add(2)

	go func(ch <-chan int,wg *sync.WaitGroup) {
      
		val,isChannelOpe := <-myCh

		//fmt.Println(<-myCh)
		fmt.Println(isChannelOpe)
        fmt.Println(val)		
		wg.Done()
	}(myCh,wg)

	go func(ch chan<- int,wg *sync.WaitGroup) { 
		myCh <- 0
		close(myCh)
		
		// myCh <- 6
		
      wg.Done()
	}(myCh,wg)

	wg.Wait()

}
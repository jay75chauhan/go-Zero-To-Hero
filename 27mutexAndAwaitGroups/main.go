package main

import (
	"fmt"
	"sync"
)

func main() {
  fmt.Println("Race Condition")

  wg:=&sync.WaitGroup{}

  mut :=&sync.RWMutex{}
  
   
  var score = []int{0}
  
  wg.Add(4)
  go func (wg *sync.WaitGroup, m *sync.RWMutex)   {
	   fmt.Println("one R")

	   mut.Lock()
	  score = append(score,1)

	   mut.Unlock()
	  wg.Done()
  }(wg,mut)
//    wg.Add(1)
   go func (wg *sync.WaitGroup,m *sync.RWMutex)  {
	   fmt.Println("two R")
	    mut.Lock()
	  score = append(score,2)
	  mut.Unlock()
	  wg.Done()
  }(wg,mut)
   go func (wg *sync.WaitGroup,m *sync.RWMutex)  {
	   fmt.Println("three R")
	    mut.Lock()
	  score = append(score,3)
	  mut.Unlock()
	  wg.Done()
  }(wg,mut)
   go func (wg *sync.WaitGroup,m *sync.RWMutex)  {
	   fmt.Println("four R")
	    mut.Lock()
	  score = append(score,4)
	  mut.Unlock()
	  wg.Done()
  }(wg,mut)
   go func (wg *sync.WaitGroup,m *sync.RWMutex)  {
	   fmt.Println("ALL R")
	   mut.RLock()
	  fmt.Println(score)
	  mut.RUnlock()

	  //if you are reading use mut.Rlock()
	  wg.Done()
  }(wg,mut)

  wg.Wait()

  fmt.Println(score)
 
}





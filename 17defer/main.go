package main

import "fmt"

func main() {
 defer	fmt.Println("hooo")
 defer	fmt.Println("hooo 55")
	 fmt.Println("ho 2")

	 myDefer()
}

func myDefer() {
	for i := 0; i < 10;i++ {
		fmt.Println(i)
	}
}
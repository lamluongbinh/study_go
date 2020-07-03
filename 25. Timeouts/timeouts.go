package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Func1"
	}()
	select {
	case res1 := <-c1:
		fmt.Println(res1)
	case <-time.After(1 * time.Second):
		fmt.Println("Func1 timeout")
	}
	//
	c2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Func2"
	}()
	select {
	case res2 := <-c2:
		fmt.Println(res2)
	case <-time.After(3 * time.Second):
		fmt.Println("Func2 timeout")
	}
}

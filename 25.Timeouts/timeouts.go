package main

import (
	"fmt"
	"time"
)

func main() {
	// Run a gorountine that will message Func1 after 2 Seconds
	c1 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Func1"
	}()
	// Select block a time < timeouts time 'in this case 1 Second'
	// So this select cannot wait to the time c1 have data -> timeouts
	// And the result will be Func1 timeout
	select {
	case res1 := <-c1:
		fmt.Println(res1)
	case <-time.After(1 * time.Second):
		fmt.Println("Func1 timeout")
	}
	// In this case select will message Func2 after 2 Seconds too
	c2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Func2"
	}()
	// But this select is set to wait for 3 Seconds before timeouts
	// So it can wait to res2 get data and then the select is success
	// Result return us c2 message: Func2å
	select {
	case res2 := <-c2:
		fmt.Println(res2)
	case <-time.After(3 * time.Second):
		fmt.Println("Func2 timeout")
	}
}

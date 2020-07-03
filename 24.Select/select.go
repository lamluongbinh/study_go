package main

import (
	"fmt"
	"time"
)

func reveive(ti int, c chan string, message string) {
	time.Sleep(time.Duration(ti) * time.Second)
	c <- message
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	ms1 := "First"
	ms2 := "Second"
	i1 := 1
	i2 := 2
	go reveive(i1, c1, ms1)
	go reveive(i2, c2, ms2)
	// A select without default will block code until a case statement is executed
	// So with default Nothing happen will be printed
	select {
	case msg1 := <-c1:
		fmt.Println("Received", msg1)
	case msg2 := <-c2:
		fmt.Println("Received", msg2)
	default:
		fmt.Println("Nothing happen")
	}
	// Without for, select will just happen 1 time, when a message arrived to msg
	// So we must add for 2 times for received all
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		}
	}
}

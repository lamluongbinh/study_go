package main

import (
	"fmt"
)

func sendstring(st string, c chan string) {
	c <- st
}

func main() {
	/***
	A small note that if We use channel without gorountine like this small code
	---
	cr := make(chan string)
	cr<-"1st"
	fmt.Println(<-cr)
	---
	We will be in a deadlock, because channel will block all the code until it have a receiver to receive message in pipeline
	So we must use a gorountine when using channel like this below code:
	***/
	cr := make(chan string)
	st1 := "1st"
	go sendstring(st1, cr)
	fmt.Println(<-cr)
	// Buffered channel will help us to store 2 values in sending pipeline
	ch := make(chan string, 2)
	ch <- "First"
	ch <- "Second"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	/***
	So if we just send 2 value to channel like code above, the code will run normaly
	But if the 3rd value send to channel, it's will cause deadlock again.
	***/
}

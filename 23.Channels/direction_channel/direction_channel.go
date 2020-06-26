package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	// We can define channel that just send (ping channel)
	// Or Receive(pong channel)
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	// Use ping to get messages
	// Then pass the message to pongs
	// and then prints message out
	ping(pings, "This is the ping pong ball, lol")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

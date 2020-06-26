package main

import (
	"fmt"
)

func sum(ar []int, c chan int) {
	sum := 0
	for _, i := range ar {
		sum += i
	}
	c <- sum
}

func main() {
	// Channel is pipeline that we can send/receive value
	// In default, channel will be unbuffered, mean that at one time just can send 1 value
	// When <-c run, the code will be block until we have some data is received by c
	ar := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int)
	fmt.Println("Split s to 3 slice and Sum it ")
	go sum(ar[:len(ar)/3], c)
	go sum(ar[len(ar)/3:len(ar)/3*2], c)
	go sum(ar[len(ar)/3*2:], c)
	x, y, z := <-c, <-c, <-c
	close(c)
	// We can see that: x,y,z will be different every time we run this example
	// Because which gorountine ended first, which will be assign first
	fmt.Println(x, y, z)
}

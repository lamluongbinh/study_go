package main

import (
	"fmt"
	"time"
)

func goroutines(st string, timeV time.Duration) {
	for i := 0; i < 100; i++ {
		fmt.Println(st, ": ", i)
		time.Sleep(timeV)
	}
}

func main() {
	// Example for 2 go process running in different threads
	var go_duration time.Duration = 400 * time.Millisecond
	var routine_duration time.Duration = 200 * time.Millisecond
	go goroutines("Go", go_duration)
	go goroutines("Routines", routine_duration)
	time.Sleep(5 * time.Minute)
	fmt.Println("Done:")
}

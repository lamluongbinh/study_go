package main

import "fmt"

func main() {
	fmt.Println("1. Simple example of switch:")
	var a = "start"
	switch a {
	case "start":
		fmt.Println("Starting programes")
	case "stop":
		fmt.Println("Stopping programes")
	}
	fmt.Println("2. Switch as a function:")
	switchfunc := func(i string) {
		switch i {
		case "start":
			fmt.Println("Starting programes")
		case "stop":
			fmt.Println("Stopping programes")
		default:
			fmt.Println(i + " is not declared.")
		}
	}
	switchfunc("start")
	switchfunc("stop")
	switchfunc("status")
	fmt.Println("2.1 Another way to declare switch as a function:")
	var switchfunc1 = func(i string) {
		switch i {
		case "start":
			fmt.Println("Starting programes")
		case "stop":
			fmt.Println("Stopping programes")
		default:
			fmt.Println(i + " is not declared.")
		}
	}
	switchfunc1("start")
}

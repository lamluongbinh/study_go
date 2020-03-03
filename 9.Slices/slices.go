package main

import "fmt"

func main() {
	fmt.Println("================= Slice Basic =================")
	s := make([]string, 3)
	fmt.Println("This is the empty slice:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("Set a,b,c to slide s: ", s)
	fmt.Println("Get data from position 1 of slice s:", s[1])
	fmt.Println("Get the length of slices:", len(s))
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("Append d,e,f to slice s:", s)

	fmt.Println("================= Slice more advance =================")
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copy slice from s to c", c)

	l := s[2:5]
	fmt.Println("Get data of s from position 2 to 5, exclude 5", l)

	l = s[:5]
	fmt.Println("Get data of s from begin to 5, exclude 5", l)

	l = s[2:]
	fmt.Println("Get data of s from 2 to end, include 2", l)

	another := []string{"I", "We", "You"}
	fmt.Println("Another way to declare slice:", another)

	fmt.Println("================= Multi demension slice =================")
	multislice := make([][]int, 3)
	fmt.Println("We can define different length for inner slice, for example:")
	multislice[0] = make([]int, 2)
	multislice[0][0] = 0
	multislice[0][1] = 1
	multislice[1] = make([]int, 1)
	multislice[1][0] = 2
	multislice[2] = make([]int, 3)
	multislice[2][0] = 3
	multislice[2][1] = 4
	multislice[2][2] = 5
	fmt.Println("Finally we get:", multislice)
}

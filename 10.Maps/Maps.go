package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["key1"] = 10
	m["key2"] = 20
	fmt.Println(m)
	value1 := m["key1"]
	fmt.Println("Get value from key1: ", value1)
	fmt.Println("Size of map:", len(m))

	delete(m, "key1")
	fmt.Println("Remove key 1 from map:", m)

	fmt.Printf("Checking if key is exist in map:")
	_, prs := m["key1"]
	fmt.Println("You can see when key1 is removed:", prs)
	_, prs = m["key2"]
	fmt.Println("And key2 is existing:", prs)

	m2 := map[string]int{"m2key1": 100, "m2key2": 200}
	fmt.Println("Another way to declare map:", m2)
}

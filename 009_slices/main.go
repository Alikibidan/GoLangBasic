package main

import "fmt"

func main() {
	var map1 = map[int]string{}
	map1[1] = "one"
	fmt.Println(map1)

	map2 := make(map[string]string)
	map2["one"] = "1"
	fmt.Println(map2)

}

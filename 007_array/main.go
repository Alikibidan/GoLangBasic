package main

import "fmt"

func main() {
	var array [5]int

	array[0] = 1111
	array[1] = 22
	array[2] = 11
	array[3] = 1
	for i, v := range array {
		fmt.Printf("Ind: %v, Val: %v\n", i, v)
	}

	var array2 = [...]string{"tt", "ff", "ss"}

	fmt.Println(len(array2))
}

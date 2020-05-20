package main

import "fmt"

func main() {
	fmt.Println("xixi")
	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)
	fmt.Println(len(arr1))
	for i, v := range arr1 {
		println(i, v)
	}
}

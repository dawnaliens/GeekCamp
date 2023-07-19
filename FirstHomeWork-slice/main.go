package main

import "fmt"

func DelSliceIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	res := DelSliceIndex(slice, 2)
	fmt.Println(res)
}

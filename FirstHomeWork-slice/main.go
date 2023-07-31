package main

import "fmt"

// 第一种方法
func DelSliceIndex(s []int, index int) []int {
	if index < 0 || index >= len(s) {
		return s
	}
	return append(s[:index], s[index+1:]...)
}

// 2. 第二种方法
// 将要删除的元素向前移动一位来覆盖，不需要重新分配内存
func DeleteSliceElement(s []int, index int) []int {
	copy(s[index:], s[index+1:])
	return s[:len(s)-1]
}

//3.目前我还不怎么理解泛型，还需要多写一下

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("The array is: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))

	slice2 := DeleteSliceElement(slice, 5)
	fmt.Println(slice2)
	fmt.Printf("The array is: %v, len: %d, cap: %d\n", slice, len(slice2), cap(slice2))
}

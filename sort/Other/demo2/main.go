package main

import (
	"fmt"
	"sort"
)

// 在排序切片中搜索
func main() {
	ints := []int{1, 3, 5, 7, 9}
	x := 5

	// 搜索前必须先排序(这里的ints是已经经过排序的)
	index := sort.SearchInts(ints, x)
	fmt.Printf("%d found at index %d\n", x, index) // 5 found at index 2

	// 对于自定义条件
	index = sort.Search(len(ints), func(i int) bool {
		return ints[i] >= x
	})
	fmt.Printf("%d found at index %d\n", x, index) // 5 found at index 2
}

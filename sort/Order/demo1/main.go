package main

import (
	"fmt"
	"sort"
)

func main() {

	// 对整数切片排序
	ints := []int{3, 1, 4, 1, 5, 9, 2, 6}
	// sort.Ints(ints)   // 升序排序
	sort.Sort(sort.IntSlice(ints))
	fmt.Println(ints) // [1 1 2 3 4 5 6 9]

	// 对浮点数切片排序
	floats := []float64{3.14, 1.41, 2.71, 0.99}
	sort.Float64s(floats)
	fmt.Println(floats) // [0.99 1.41 2.71 3.14]

	// 对字符串切片排序
	strings := []string{"banana", "apple", "cherry"}
	sort.Strings(strings)
	fmt.Println(strings) // ["apple", "banana", "cherry"]
}

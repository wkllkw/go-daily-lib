package main

import (
	"fmt"
	"sort"
)

func main() {
	// 对整数切片逆序排序
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println(ints)

	// 对浮点数切片逆序排序
	floats := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7}
	sort.Sort(sort.Reverse(sort.Float64Slice(floats)))
	fmt.Println(floats)

	// 对字符串切片逆序排序
	strings := []string{"a", "b", "c", "d", "e", "f", "g"}
	sort.Sort(sort.Reverse(sort.StringSlice(strings)))
	fmt.Println(strings)
}

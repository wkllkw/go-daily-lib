package main

import (
	"fmt"
	"sort"
)

// 检查切片是否已排序
func main() {
	ints := []int{1, 2, 3, 4}
	fmt.Println(sort.IntsAreSorted(ints)) // true

	ints = []int{4, 3, 2, 1}
	fmt.Println(sort.IntsAreSorted(ints)) // false
}

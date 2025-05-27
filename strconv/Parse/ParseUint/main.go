package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := "42"
	// 第二个参数表示字符串的进制基数（base）
	if s, err := strconv.ParseUint(v, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	// 第三个参数表示结果的位数（32 表示解析为 Unit32 类型，但返回值仍是 Unit64 类型，只是会检查是否在 Unit32 范围内）。
	if s, err := strconv.ParseUint(v, 10, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
}

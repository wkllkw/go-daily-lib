package main

import (
	"fmt"
	"strconv"
)

func main() {
	v32 := "-354634382"
	// 第二个参数表示字符串的进制基数（base）
	if s, err := strconv.ParseInt(v32, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	// 第三个参数表示结果的位数（32 表示解析为 int32 类型，但返回值仍是 int64 类型，只是会检查是否在 int32 范围内）。
	if s, err := strconv.ParseInt(v32, 16, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

	v64 := "-3546343826724305832"
	if s, err := strconv.ParseInt(v64, 10, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseInt(v64, 16, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
}

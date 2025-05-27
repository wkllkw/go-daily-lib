package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := uint64(42)

	/*
		第二个参数指定转换的进制基数（取值范围：2 到 36）。
		10：十进制（输出如 "-42"）。
		2：二进制（输出如 "-101010"）。
		16：十六进制（输出如 "-2a"）。
		其他：支持任意进制（如 36 会用 0-9a-z 表示）。
	*/
	s10 := strconv.FormatUint(v, 10)
	fmt.Printf("%T, %v\n", s10, s10)

	s16 := strconv.FormatUint(v, 16)
	fmt.Printf("%T, %v\n", s16, s16)
}

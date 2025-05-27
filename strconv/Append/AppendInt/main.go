package main

import (
	"fmt"
	"strconv"
)

func main() {
	b10 := []byte("int (base 10):")

	/*
		第一个参数是目标字节切片，函数会将整数转换后的字符串 追加到这个切片的末尾。
		第二个参数要转换为字符串的整数值（可以是正数、负数或零）。
		第三个参数指定转换的 进制基数，取值范围为 2 到 36
	*/
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10))

	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16))
}

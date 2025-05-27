package main

import (
	"fmt"
	"strconv"
)

func main() {
	// IsPrint 报告该符文是否被 Go 定义为可打印，其定义与 unicode.IsPrint 相同：字母、数字、标点符号、符号和 ASCII 空格。
	c := strconv.IsPrint('\u263a')
	fmt.Println(c)

	bel := strconv.IsPrint('\007')
	fmt.Println(bel)
}

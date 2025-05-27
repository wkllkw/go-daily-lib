package main

import (
	"fmt"
	"strconv"
)

func main() {
	// CanBackquote 报告字符串 s 是否可以不变地表示为单行反引号字符串，其中不包含制表符以外的控制字符。
	fmt.Println(strconv.CanBackquote("Fran & Freddie's Diner ☺"))
	fmt.Println(strconv.CanBackquote("`can't backquote this`"))
}

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// 格式化的模板为 2006-01-02 15:04:05

	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))

	// 小数点后写0，因为有3个0所以格式化输出的结果也保留3位小数
	fmt.Println(now.Format("2006/01/02 15:04:05.000"))
	// 小数点后写9，会省略末尾可能出现的0
	fmt.Println(now.Format("2006/01/02 15:04:05.999"))

	// 只格式化时分秒部分
	fmt.Println(now.Format("15:04:05"))
	// 只格式化日期部分
	fmt.Println(now.Format("2006.01.02"))
}

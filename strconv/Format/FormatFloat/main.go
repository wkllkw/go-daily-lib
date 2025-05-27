package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := 3.1415926535

	// 第二个参数指定输出格式，第三个参数控制小数位数，第三个参数指定来源浮点数的位数
	/*
		标记		说明									示例输入			示例输出
		'f'		普通十进制格式							123.456			"123.456"
		'E'		科学计数法（大写 E）					123.456			"1.23456E+02"
		'e'		科学计数法（小写 e）					123.456			"1.23456e+02"
		'g'		自动选择 'f' 或 'e'（更紧凑的格式）		123.456			"123.456"
		'b'		指数为 2 的科学计数法（罕见用途）			123.456			"8684392467609p-46"
	*/
	s32 := strconv.FormatFloat(v, 'E', -1, 32)
	fmt.Printf("%T, %v\n", s32, s32)

	s64 := strconv.FormatFloat(v, 'E', -1, 64)
	fmt.Printf("%T, %v\n", s64, s64)

	// fmt.Println uses these arguments to print floats
	fmt64 := strconv.FormatFloat(v, 'g', -1, 64)
	fmt.Printf("%T, %v\n", fmt64, fmt64)
}

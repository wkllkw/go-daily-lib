package main

import (
	"fmt"
	"strconv"
)

func main() {
	b32 := []byte("float32:")

	/*
		第一个参数目标字节切片（[]byte），函数会将浮点数转换后的字符串 追加 到这个切片末尾。
		第二个参数要转换为字符串的浮点数值
		第三个参数指定浮点数的输出格式
		第四个参数指定浮点数的精度，-1指的是默认精度，删除无意义的0
		第五个参数指定浮点数的原始位数
	*/
	b32 = strconv.AppendFloat(b32, 3.1415926535, 'E', -1, 32)
	fmt.Println(string(b32))

	b64 := []byte("float64:")
	b64 = strconv.AppendFloat(b64, 3.1415926535, 'E', -1, 64)
	fmt.Println(string(b64))
}

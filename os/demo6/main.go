package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取命令行参数
	args := os.Args
	fmt.Println("Program:", args[0])
	fmt.Println("Arguments:", args[1:])
}

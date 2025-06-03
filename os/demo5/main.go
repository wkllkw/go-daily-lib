package main

import (
	"fmt"
	"os"
)

func main() {
	// 设置环境变量
	os.Setenv("NAME", "gopher")

	// 获取环境变量
	name := os.Getenv("NAME")
	fmt.Println(name)

	// 获取所有环境变量
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	// 查找环境变量
	path, found := os.LookupEnv("PATH")
	if found {
		fmt.Println("PATH:", path)
	}
}

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 获取当前工作目录
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Working directory:", wd)

	// 改变工作目录
	err = os.Chdir(".\\tmp")
	if err != nil {
		log.Fatal(err)
	}
}

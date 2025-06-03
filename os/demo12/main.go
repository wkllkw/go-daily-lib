package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 创建符号链接
	err := os.Symlink("original.txt", "link.txt")
	if err != nil {
		log.Fatal(err)
	}

	// 读取符号链接
	target, err := os.Readlink("link.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Link points to:", target)
}

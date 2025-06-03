package main

import (
	"log"
	"os"
)

func main() {
	// 修改文件权限
	err := os.Chmod("file.txt", 0644)
	if err != nil {
		log.Fatal(err)
	}

	// 修改文件所有者
	err = os.Chown("file.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Fatal(err)
	}
}

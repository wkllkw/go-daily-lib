package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 获取文件信息
	fileInfo, err := os.Stat("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size:", fileInfo.Size(), "bytes")
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is directory:", fileInfo.IsDir())
}

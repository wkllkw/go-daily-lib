package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	OpenFileRW()
}

func OpenFile() {
	// 打开文件(只读)
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 读取文件内容
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read %d bytes: %q\n", count, data[:count])
}

func OpenFileRW() {
	// 以读写方式打开文件
	file, err := os.OpenFile("file.txt", os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()

	// 写入文件
	_, err = file.WriteString("Hello, World!\n")
	if err != nil {
		log.Fatal(err)
	}
}

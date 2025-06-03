package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 读取目录内容
	files, err := os.ReadDir("./test")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

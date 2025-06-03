package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	CreateFileTemp()
}

func CreateFileTemp() {
	// 创建临时文件
	tempFile, err := os.CreateTemp(".", "example.*.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			log.Fatal(err)
		}
	}(tempFile.Name()) // 清理
	fmt.Println("Temp file:", tempFile.Name())
}

func CreateDirTemp() {
	// 创建临时目录
	tempDir, err := os.MkdirTemp(".", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			log.Fatal(err)
		}
	}(tempDir) // 清理
	fmt.Println("Temp dir:", tempDir)
}

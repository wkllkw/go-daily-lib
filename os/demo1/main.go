package main

import (
	"log"
	"os"
)

func main() {
	Remove()
}
func CreateFile() {
	// 创建文件
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

func Mkdir() {
	// 创建目录
	err := os.Mkdir("mydir", 0755) // 0755是权限模式
	if err != nil {
		log.Fatal(err)
	}
}

func MkdirAll() {
	// 创建多级目录
	err := os.MkdirAll("path/to/dir", 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func Remove() {
	// 删除文件
	err := os.Remove("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	// 删除目录(必须为空)
	err = os.Remove("mydir")
	if err != nil {
		log.Fatal(err)
	}
}

func RemoveAll() {
	// 递归删除目录
	err := os.RemoveAll("path/to/dir")
	if err != nil {
		log.Fatal(err)
	}
}

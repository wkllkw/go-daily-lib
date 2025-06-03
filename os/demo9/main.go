package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 获取当前用户
	user, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Home directory:", user)

	// 获取用户ID和组ID
	fmt.Println("User ID:", os.Getuid())
	fmt.Println("Group ID:", os.Getgid())
}

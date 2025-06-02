package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

// 自定义结构体的排序
func main() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	// 按年龄排序
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people) // [{Michael 17} {Jenny 26} {Bob 31} {John 42}]

	// 按姓名排序
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	fmt.Println(people) // [{Bob 31} {Jenny 26} {John 42} {Michael 17}]
}

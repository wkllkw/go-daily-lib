package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		Unquote 将 s 解释为单引号、双引号或反引号的 Go 字符串字面量，并返回 s 所引用的字符串值。
		（如果 s 是单引号，则为 Go 字符字面量；Unquote 返回相应的单字符字符串。如果为空字符字面量，Unquote 返回空字符串。）
	*/
	s, err := strconv.Unquote("You can't unquote a string without quotes")
	fmt.Printf("%q, %v\n", s, err)
	s, err = strconv.Unquote("\"The string must be either double-quoted\"")
	fmt.Printf("%q, %v\n", s, err)
	s, err = strconv.Unquote("`or backquoted.`")
	fmt.Printf("%q, %v\n", s, err)
	s, err = strconv.Unquote("'\u263a'") // single character only allowed in single quotes
	fmt.Printf("%q, %v\n", s, err)
	s, err = strconv.Unquote("'\u2639\u2639'")
	fmt.Printf("%q, %v\n", s, err)
}

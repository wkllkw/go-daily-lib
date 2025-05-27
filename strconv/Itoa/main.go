package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a = 1000

	b := strconv.Itoa(a)

	fmt.Printf("%#v", b)

}

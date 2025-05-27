package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a = "1000"

	b, _ := strconv.Atoi(a)

	fmt.Printf("%+v", b)

}

package main

import "fmt"

func main() {
	var first_var int64 = 0xA
	var second_var int = int(first_var)

	fmt.Println(second_var)
}

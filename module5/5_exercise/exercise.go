package main

import "fmt"

func main() {
	const (
		_ = iota
		januaryMonth
		febraryMonth
		marchMonth
		aprilMonth
		mayMonth
	)

	fmt.Println(januaryMonth, febraryMonth, marchMonth, aprilMonth, mayMonth)
}

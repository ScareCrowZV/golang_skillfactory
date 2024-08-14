package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	for j := 0; j < 10; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Println(j)
	}

	var number = [10]int{1. - 2, 3, -4, 5, -6, -7, 8, -9, 10}
	var sum = 0

	for _, value := range number {
		if value < 0 {
			continue
		}
		sum += value
	}
	fmt.Println("Sum:", sum)

outerLoop:
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if i*j > 100 {
				break outerLoop
			}
		}
	}

	a, b, c, d, e := 3, 4, 5, 6, 7

	if a == 3 {
		if b == 4 {
			if c == 5 {
				if d == 6 {
					if e == 7 {
						fmt.Println("Цикломатическая сложность: 5")
					}
				}
			}
		}
	}

	// for {
	// 	switch {
	// 	default:
	// 		break
	// 	}
	// 	break

	// }

	number20 := 0

	for i := 20; i > 0; i-- {
		number20 = i
	}

	fmt.Println(number20)

	for i := 1; i < 10; i *= 2 {
		fmt.Println(i)
	}

}

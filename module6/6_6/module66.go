package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	// var text2 string
	var number int
	fmt.Print("Введите число: ")
	_, err := fmt.Scanln(&number)
	if err != nil {
		fmt.Println(fmt.Sprintf("Не получилось прочитать строку: %s", err))
	}

	// number, err := strconv.Atoi(text2)
	// if err != nil {
	// 	fmt.Println(fmt.Sprintf("Не получилось сконвертировать строку в число: %s", err))
	// }

	if err == nil {
		fmt.Print(number)
	}
}

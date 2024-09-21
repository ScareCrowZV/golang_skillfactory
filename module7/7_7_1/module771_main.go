package main

import (
	"bufio"
	"calc"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var expression string
	var err error
	var reader *bufio.Reader
	var firstOperand float64
	var firstOperandBuffer string
	var firstOperandBufferFull bool = false
	var secondOperand float64
	var secondOperandBuffer string
	var operator byte
	var result float64
	var errorResult string

	fmt.Print(
		"Вас приветствует простой калькулятор!\n",
		"Введите выражение и нажмите клавишу Enter\n",
		"Например: \"1 + 1\" Enter\n",
		"И даже вещественные числа: \"10.5 \\ 5.5\" Enter\n",
		"Учитываться будут только первые 2 операнда\n",
		"\nВозможные операции:\n",
		"+ сложение\n",
		"- вычитание\n",
		"* умножение\n",
		"\\ деление\n")

	reader = bufio.NewReader(os.Stdin)
	expression, err = reader.ReadString('\n')

	if err != nil {
		fmt.Printf("Ошибка ввода: %s", err)
		return
	}

	for i := 0; i < len(expression); i++ {

		switch {
		case (expression[i] >= 48 && expression[i] <= 57 || string(expression[i]) == "." || string(expression[i]) == ",") && firstOperandBufferFull == false:
			firstOperandBuffer += string(expression[i])
		case (expression[i] >= 48 && expression[i] <= 57 || string(expression[i]) == "." || string(expression[i]) == ",") && firstOperandBufferFull == true && operator != 0:
			secondOperandBuffer += string(expression[i])
		case string(expression[i]) == "+":
			operator = expression[i]
			firstOperandBufferFull = true
		case string(expression[i]) == "-":
			operator = expression[i]
			firstOperandBufferFull = true
		case string(expression[i]) == "*":
			operator = expression[i]
			firstOperandBufferFull = true
		case string(expression[i]) == "\\":
			operator = expression[i]
			firstOperandBufferFull = true
		case expression[i] == 13: // Перенос строки игнорируем
		case expression[i] == 10: // Возврат каретки игнорируем
		case expression[i] == 32: // Пробел игнорируем
		default:
			fmt.Printf("Введён недопустимый символ %s, выражение введено не верно\n", string(expression[i]))
			return
		}
	}

	// Пробуем привести первый операнд
	firstOperand, err = strconv.ParseFloat(firstOperandBuffer, 64)
	secondOperand, err = strconv.ParseFloat(secondOperandBuffer, 64)

	cc := calc.NewCalculator()
	result, errorResult = cc.Calculate(firstOperand, secondOperand, operator)

	if errorResult != "" {
		fmt.Printf("Во время вычисления произошла ошибка: %s\n", errorResult)
		return
	}

	if operator > 0 {
		fmt.Printf("Ваше выражение: %f %s %f = %f\n", firstOperand, string(operator), secondOperand, result)
	}

	fmt.Println("Нажмите Enter для завершения")
	_, _ = reader.ReadString('\n')

}

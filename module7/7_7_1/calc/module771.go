package calc

type calculator struct {
}

func NewCalculator() calculator {
	return calculator{}
}

// Сложение
func (c calculator) addition(firstOperand, secondOperand float64) float64 {
	return firstOperand + secondOperand
}

// Вычитание
func (c calculator) subtraction(firstOperand, secondOperand float64) float64 {
	return firstOperand - secondOperand
}

// Умножение
func (c calculator) multiplication(firstOperand, secondOperand float64) float64 {
	return firstOperand * secondOperand
}

// Деление
func (c calculator) division(firstOperand, secondOperand float64) (float64, string) {
	if secondOperand == 0 {
		return 0, "Деление на 0 не возможно"
	}
	return firstOperand / secondOperand, ""
}

func (c calculator) Calculate(firstOperand, secondOperand float64, operator byte) (float64, string) {

	var result float64
	var errorResult string

	switch string(operator) {
	case "+":
		result = c.addition(firstOperand, secondOperand)
		break
	case "-":
		result = c.subtraction(firstOperand, secondOperand)
		break
	case "*":
		result = c.multiplication(firstOperand, secondOperand)
		break
	case "\\":
		result, errorResult = c.division(firstOperand, secondOperand)
		break
	default:
		errorResult = "Выражение введено не верно"
	}

	return result, errorResult

}

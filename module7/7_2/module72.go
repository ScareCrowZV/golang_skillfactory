package main

import "fmt"

var (
	ErrZeroA       = fmt.Errorf("coefficient a is zero")      // уравнение не является квадратным
	ErrNoRealRoots = fmt.Errorf("equation has no real roots") // у уравнения нет вещественных корней
)

func main() {

}

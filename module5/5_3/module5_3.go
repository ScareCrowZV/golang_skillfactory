package main

import "fmt"

func main() {
	fmt.Println("Привет, всем!")
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var fl32 float32
	var fl64 float64
	var cplx64 complex64
	var cplx128 complex128
	var rn rune
	var str string
	var is bool

	fmt.Println(i8, i16, i32, i64, fl32, fl64, cplx64, cplx128, rn, str, is)

	i8 = -128

	fmt.Println(i8)

	i8 = 0b01100100

	fmt.Println(i8)

	i8 = 0o070

	fmt.Println(i8)

	i8 = 0xF

	fmt.Println(i8)

	varshort := 100

	fmt.Println(varshort)

	var many_a, many_b, many_c int

	fmt.Println(many_a, many_b, many_c)

	m_a, m_b, m_c := 0, 1, 2

	fmt.Println(m_a, m_b, m_c)

	var (
		ma_a int
		ma_b int32
		ma_c uint64
	)

	fmt.Println(ma_a, ma_b, ma_c)

	var realNumber float32

	fmt.Println(realNumber)

	realNumberShort := 100 // Это будет int

	fmt.Println(realNumberShort)

	realRealNumberShort := 100.0 // это будет float

	fmt.Println(realRealNumberShort)

	realHexExp := 0x2.p10
	realHexExp2 := 0x.8p-0
	realDecExp := 6.67428e-11
	realDecExp2 := 1e6
	realDec := .25

	fmt.Println(realHexExp, realHexExp2, realDecExp, realDecExp2, realDec)

	var symbol rune
	symbol2 := 'Я'
	var symbol3 rune = '\u2626'

	fmt.Println(symbol, symbol2, string(symbol3))

	var stringU string = `lopata
	100500
	master`

	var stringUU string = "\nhi! \"\\ \077"
	fmt.Println(stringU, stringUU)
	fmt.Println(stringUU[5])

	t := true
	f := false
	fmt.Println(t)
	fmt.Println(f)

	const pi = 3.14
	const piT float64 = 3.14
	const piPlusOne = piT + 1
	const (
		constA int = 5
		constB     = 4e1
	)

	fmt.Println(pi, piT, piPlusOne, constA, constB)

	const (
		c_iota_0 = iota
		c_iota_1 = iota
		c_iota_2 = iota
	)

	fmt.Println(c_iota_0, c_iota_1, c_iota_2)

	const (
		c_iota_3 = 50 + iota
		c_iota_4 = c_iota_3 + iota
		c_iota_5 = 3
		c_iota_6 = 50 + iota
	)

	fmt.Println(c_iota_3, c_iota_4, c_iota_5, c_iota_6)

	const (
		c_iota_7         = iota * 42
		c_iota_8 float64 = iota * 42
		c_iota_9         = iota * 42
	)

	fmt.Println(c_iota_7, c_iota_8, c_iota_9)

	const (
		c_iota_10 = iota
		_         = iota
		c_iota_11 = iota
	)

	fmt.Println(c_iota_10, c_iota_11)

	const (
		c_iota_12 = iota
		_         = iota
		c_iota_13 = iota
	)

	fmt.Println(c_iota_12, c_iota_13)
}

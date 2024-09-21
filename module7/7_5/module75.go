/*
Помимо новых структур, с помощью ключевого слова type можно объявлять типы соответствующие встроенным типам GoLang
*/
package main

func main() {

	type MyInt int64
	// Работает с копией структуры
	// func (c Circle) Area() float64 { /* ... */ }

	// Работает с указателем на структуру
	// func (c *Circle) Area() float64 { /* ... */ }

	// Работает с
	// func (c *Circle) Area() float64 {
	// 	return math.Pi * c.r * c.r
	// }

}

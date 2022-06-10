// Package divide_doc implements functions to demonstrate godoc features
// The divide_Doc returns some result from the division of two integers
// If one of the numbers is zero, an error will be occurred
// if the result of the division is zero, then a message will be displayed

package divide_Doc

import (
	"fmt"
	"log"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	safeDivide(a, b)
}

func safeDivide(a, b int) {

	defer func() {
		if err := recover(); err != nil {
			if fmt.Sprintf("%v", err) == "runtime error: integer darwinDivide by zero" {
				log.Println("Введите делимое не нулевое значение")
				return
			}
			log.Printf("Panic: %v", err)
		}
	}()

	fmt.Println("Результат: ", a/b)
}

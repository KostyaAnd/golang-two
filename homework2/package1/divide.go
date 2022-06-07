package package1

import (
	"fmt"
	"log"
)

func Divide() {
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

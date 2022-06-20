package main

import (
	"fmt"
)

var thousandGoroutines = 1000

func main() {
	ch := make(chan int, 1000)

	for i := 0; i < thousandGoroutines; i++ {
		go func() {
			ch <- 1
		}()
	}

	var total int

	for i := 0; i < thousandGoroutines; i++ {
		<-ch
		total++
	}
	fmt.Println(total)
}

package main

import "fmt"

func fibonacci(cacheMap map[int]int, n int) int {
	if n < 2 {
		return n
	}

	f, ok := cacheMap[n-1]
	if !ok {
		f = fibonacci(cacheMap, n-1)
		cacheMap[n-1] = f
	}

	return fibonacci(cacheMap, n-2) + f
}
func main() {
	var j int
	fmt.Println("Введите число:")
	fmt.Scanln(&j)
	cacheMap := make(map[int]int, 20)
	for i := 0; i < j; i++ {
		fmt.Printf("%d\t", fibonacci(cacheMap, i))
	}
}

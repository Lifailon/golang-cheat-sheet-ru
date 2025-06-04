package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	array := make([]int, N)
	for i := 0; i < N; i++ {
		var a int
		fmt.Scan(&a)
		array[i] = a
	}
	var count int
	for _, e := range array {
		if e >= 0 {
			count++
		}
	}
	fmt.Print(count)
}

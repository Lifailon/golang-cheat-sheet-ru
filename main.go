package main

import "fmt"

func main() {
	// Считываем в формате строки
	var left string
	fmt.Scan(&left)
	var right string
	fmt.Scan(&right)
	// Разбиваем строку на массив из букв (rune)
	for _, l := range left {
		for _, r := range right {
			if l == r {
				fmt.Print(int(l) - '0')
				fmt.Print(" ")
				break
			}
		}
	}
}

package main

import (
	"fmt"
)

func main() {
	var i int
main:
	for {
		fmt.Scan(&i)
		switch {
		case i < 10:
			continue main
		case i > 100:
			break main
		default:
			fmt.Println(i)
		}
	}
}

package main

import "fmt"

func main() {
	var v int = 1
	{
		fmt.Println(v) // 1
		var v string = "2"
		fmt.Println(v) // 2
	}
	fmt.Println(v) // 1
}

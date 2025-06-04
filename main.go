package main

import "fmt"

func main() {
	// Создаем срез длинной 10 элементов с типом данных uint8
	var workArray [10]uint8
	var counter int = 0
	// Заполняем срез из stdin по индексу счетчика counter
	for {
		if counter > 9 {
			break
		}
		var input uint8
		fmt.Scan(&input)
		workArray[counter] = input
		counter++
	}
	// Создаем цикл из трех интераций
	for i := 0; i < 3; i++ {
		// Считываем пару индексов из stdin
		var leftInputIndex int
		var rightInputIndex int
		fmt.Scan(&leftInputIndex)  // 1
		fmt.Scan(&rightInputIndex) // 2
		// Фиксируем текущие значения элементов по индексам
		leftElement := workArray[leftInputIndex]   // [1] = 151
		rightElement := workArray[rightInputIndex] // [2] = 137
		// Меняем значения элементов местами по индексам
		workArray[leftInputIndex] = rightElement // [1] = 137
		workArray[rightInputIndex] = leftElement // [2] = 151
	}
	for _, element := range workArray {
		fmt.Print(element, " ")
	}
	// Проверяем тип данных массива и завершаем программу
	if _, ok := interface{}(workArray).([10]uint8); ok {
		fmt.Print("type ok")
		return
	} else {
		fmt.Print("type error")
		return
	}
}

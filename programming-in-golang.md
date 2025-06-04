Решение задач из курса [Программирование на Golang](https://stepik.org/course/54403) от Stepik.

Оглавление:

- [Арифметические операции](#арифметические-операции)
- [Условия](#условия)
- [Циклы](#циклы)
- [Форматирование](#форматирование)
- [Массивы](#массивы)

## Арифметические операции

Напишите программу, которая последовательно делает следующие операции с введённым числом:

- Число умножается на 2
- Затем к числу прибавляется 100

Sample Input:
```
1
```

Code:
```go
package main

import "fmt"

func main(){
    var a int
    fmt.Scan(&a)
    a = a*2
    a = a +100
    fmt.Print(a)
}
```

Sample Output:
```
102
```

---

Сначала находим квадраты двух чисел, а затем суммируем их.

Sample Input:
```
2
2
```

Code:
```go
package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	a = a * a
	b = b * b
	c = a + b
	fmt.Println(c)
}
```

Sample Output:
```
8
```

---

По данному целому числу, найдите его квадрат.

Sample Input:
```
3
```

Code:
```go
package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Print(a * a)
}
```

Sample Output:
```
9
```

---

Дано натуральное число, выведите его последнюю цифру.

Sample Input:
```
123
```

Code:
```go
package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s) // 123
	last := s[len(s)-1]
	str := string(last)
	fmt.Print(str)
}
```

Sample Output:
```
3
```

---

Дано неотрицательное целое число. Найдите число десятков (то есть вторую цифру справа, или второй символ строки с конца).

Sample Input:
```
2010
```

Code:
```go
package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)
	fmt.Println(string(n[len(n)-2]))
}
```

Sample Output:
```
1
```

---

Часовая стрелка повернулась с начала суток на `d` градусов (0 < `d` < 360). Определите, сколько сейчас целых часов `h` и целых минут `m`. Циферблат часов стандартный 12-ти часовой.

Sample Input:
```
90
```

Code:
```go
package main

import "fmt"

func main() {
	var d int
	fmt.Scan(&d)
	h := d / 30
	m := d % 30
	m = int(float64(m)/0.5 + 0.5)
	fmt.Printf("It is %d hours %d minutes.\n", h, m)
}
```

Sample Output:
```
It is 3 hours 0 minutes.
```

---

## Условия

На ввод подается целое число. Если число положительное - вывести сообщение `Число положительное`, если число отрицательное - `Число отрицательное`. Если подается ноль - вывести сообщение `Ноль`.

Sample Input:
```
5
```

Code:
```go
package main

import "fmt"

func main() {
	var in int
	fmt.Scan(&in)
	switch {
	case in > 0:
		fmt.Print("Число положительное")
	case in < 0:
		fmt.Print("Число отрицательное")
	default:
		fmt.Print("Ноль")
	}
}
```

Sample Output:
```
Число положительное
```

---

По данному трехзначному числу определите, все ли его цифры различны.

Sample Input:
```
237
```

Code:
```go
package main

import "fmt"

func main() {
	var in int
	fmt.Scan(&in)
	// 2
	a := in / 100
	// 3
	b := in / 10 % 10
	// 7
	c := in % 10
	// 2 != 3 && 3 != 7 && 7 != 2
	if a != b && b != c && c != a {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
```

Sample Output:
```
YES
```

---

Дано неотрицательное целое число. Найдите и выведите первую цифру числа. 

Sample Input:
```
1234
```

Code:
```go
package main

import "fmt"

func main() {
	var in string
	fmt.Scan(&in)
	// Преобразуем из кода символа Unicode (UTF-8) в целое число
	// Вычитаем значение символа 0 (48) и получаем остаток целого числа
	fmt.Print(int(in[0] - '0'))
}
```

Sample Output:
```
1
```

---

Определите является ли билет счастливым. Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.

Sample Input:
```
613244
```

Code:
```go
package main

import "fmt"

func main() {
	var a, b, c int
	var q, w, e int
	fmt.Scanf("%1d%1d%1d%1d%1d%1d", &a, &b, &c, &q, &w, &e)
	if a+b+c == q+w+e {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
```

Sample Output:
```
YES
```

---

Требуется определить, является ли данный год високосным.

Год является високосным если он соответствует хотя бы одному из нижеперечисленных условий:

- кратен 400
- кратен 4, но не кратен 100

Sample Input:
```
2000
```

Code:
```go
package main

import "fmt"

func isLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}

func main() {
	var year int
	fmt.Scan(&year)
	if isLeapYear(year) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
```

Sample Output:
```
YES
```

---

## Циклы

Напишите программу, которая выводит квадраты натуральных чисел от 1 до 10. Квадрат каждого числа должен выводиться в новой строке.

Input:
```
1
2
3
4
5
6
7
8
9
10
```

Code:
```go
package main

import "fmt"

func main() {
	var i int = 1
	for i < 11 {
		fmt.Println(i * i)
		i++
	}
}
```

Output:
```
1
4
9
16
25
36
49
64
81
100
```

---

Требуется написать программу, при выполнении которой с клавиатуры считываются два натуральных числа A и B (каждое не более 100, A < B). Вывести сумму всех чисел от A до B  включительно.

Sample Input:
```
1 5
```

Code:
```go
package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Scan(&a, &b)
	var sum int
	for i := a; i <= b; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
}
```

Sample Output:
```
15
```

---

Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8. Программа в первой строке получает на вход число n - количество чисел в последовательности, во второй строке -- n чисел, входящих в данную последовательность.

> Входящие данные должны удовлетворять двум условиям - числов в диапазоне от 10 до 99 и при делении на 8 остаток равен нулю. Получить сумму отфильтрованных входных значений.

Sample Input:
```
5
38 24 800 8 16
```

Code:
```go
package main

import (
	"fmt"
)

type Input struct {
	Value int
}

func main() {
	var n int
	// Получаем количество чисел для цикла
	fmt.Scan(&n)
	// Массив с данными из структуры
	arr := make([]Input, n)
	// Читаем все входные значения
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i].Value)
	}
	var sum int
	for _, i := range arr {
		// Проверяем входные значения
		if i.Value > 9 && i.Value < 100 && i.Value%8 == 0 {
			sum = sum + i.Value
		}
	}
	fmt.Println(sum)
}
```

> 24 + 16

Sample Output:
```
40
```

---

Последовательность состоит из натуральных чисел и завершается числом 0. Определите количество элементов этой последовательности, которые равны ее наибольшему элементу. Вводится непустая последовательность натуральных чисел, оканчивающаяся числом 0 (само число 0 в последовательность не входит, а служит как признак ее окончания).

Sample Input:
```
1
3
3
1
0
```

Code:
```go
package main

import (
	"fmt"
)

func main() {
	// Максимальное число
	var maxInt int
	// Сумма чисел с максимальным значением
	var sum int
	for {
		var n int
		fmt.Scan(&n)
		if n == 0 {
			// Останавливаем бесконечный цикл
			break
		}
		switch {
		// Сбрасываем счетчик суммы
		case n > maxInt:
			maxInt = n
			sum = 1
		// Увеличиваем счетчик
		case n == maxInt:
			sum++
		}
	}
	fmt.Print(sum)
}
```

Sample Output:

> 3 & 3

```
2
```

---

Вывести первое число от 1 до `n` включительно, кратное числу `c` (делится на `c` без остатка), но НЕ кратное `d`. Если такого числа нет - выводить ничего не нужно.


Sample Input:
```
20
3
5
```

Code:
```go
package main

import "fmt"

func main() {
	// Счетчик для цикла
	var counter int = 1
	// Входные значения
	var n int
	fmt.Scan(&n)
	var c int
	fmt.Scan(&c)
	var d int
	fmt.Scan(&d)
	for {
		if counter%c == 0 && counter%d != 0 {
			fmt.Print(counter)
			break
		} else if counter > n {
			break
		} else {
			counter++
		}
	}
}
```

Sample Output:
```
3
```

---

Напишите программу, которая считывает целые числа с консоли по одному числу в строке.

Для каждого введённого числа проверить:

- Если число меньше 10, то пропускаем это число
- Если число больше 100, то прекращаем считывать числа
- В остальных случаях вывести это число обратно на консоль в отдельной строке

Sample Input:
```
30
11
7
101
```

Code:
```go
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
```

Sample Output:
```
30
11
```

---

Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов, после чего дробная часть копеек отбрасывается. Каждый год сумма вклада становится больше. Определите, через сколько лет вклад составит не менее y рублей.

Sample Input:
```
100
10
200
```

Code:
```go
package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)
	var p float64
	fmt.Scan(&p)
	var y float64
	fmt.Scan(&y)
	sum := x
	res := 0
	for {
		// 100 * (1 + 10/100) = 110
		sum = sum + (sum / 100 * p)
		res++
		if sum >= y {
			fmt.Println(res)
			break
		}
	}
}
```

Sample Output:
```
8
```

---

Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа. Гарантируется, что цифры в числах не повторяются. Числа в пределах от 0 до 10000. Программа должна вывести цифры, которые имеются в обоих числах, через пробел. Цифры выводятся в порядке их нахождения в первом числе.

Sample Input:
```
564 8954
```

Code:
```go
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
				// Преобразуем rune в int
				// fmt.Print(int(l) - '0')
				fmt.Printf("%c", l)
				fmt.Print(" ")
				break
			}
		}
	}
}
```

Sample Output:
```
5 4
```

## Форматирование

На вход подается число типа float64. Вам нужно вывести преобразованное число по правилу: число возводится в квадрат, затем обрезается дробная часть так что остается 4 знака после запятой. Но если число равно или меньше нуля - выводить:
"число R не подходит", где R - исходное число с 2 цифрами после запятой и с 2 по ширине. А если число больше чем 10 000 - выводить исходное число в экспоненциальном представлении (знак экспоненты в нижнем регистре).

Sample input:
```
12.12345678
```

Code:
```go
package main

import (
	"fmt"
)

func main() {
	var R float64
	fmt.Scan(&R)
	if R <= 0 {
		fmt.Printf("число %2.2f не подходит", R)
		return
	}
	if R > 10000 {
		fmt.Printf("%e", R)
		return
	}
	RR := R * R
	fmt.Printf("%.4f", RR)
}

```

Sample output:
```
146.9782
```

## Массивы

На первом этапе на стандартный ввод подается 10 целых положительных чисел, которые должны быть записаны в порядке ввода в массив из 10 элементов. Тип чисел, входящих в массив, должен соответствовать минимально возможному целому беззнаковому типу. Имя массива который вы должны сами создать `workArray` (условие обязательное). Использование массива - обязательное условие.

На втором этапе на стандартный ввод подаются еще 3 пары чисел - индексы элементов этого массива, которые требуется поменять местами (если такая пара чисел 3 и 7, значит в массиве элемент с 3 индексом нужно поменять местами с элементом, индекс которого 7).

Sample Input:
```
99 151 137 71 117 187 20 93 187 67 1 2 3 5 7 8
```

Code:
```go
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
```

Sample Output:
```
99 137 151 187 117 71 20 187 93 67 type ok
```
package exercise

import (
	"fmt"
)

func Exercise1() {
	var op string
	var a, b float64
	fmt.Println("첫 번재 숫자 입력:")
	fmt.Scan(&a)
	fmt.Println("두 번째 숫자 입력:")
	fmt.Scan(&b)
	fmt.Println("연산자 입력(+,-,*,/):")
	fmt.Scan(&op)

	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("error")
		} else {
			fmt.Println(a / b)
		}
	}

}

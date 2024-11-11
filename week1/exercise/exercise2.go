package exercise

import "fmt"

func Exercise2() {
	var temp int
	fmt.Println("숫자를 입력하세요: ")
	fmt.Scan(&temp)
	if temp%2 == 0 {
		fmt.Println("짝수입니다.")
	} else {
		fmt.Println("홀수입니다.")
	}
}

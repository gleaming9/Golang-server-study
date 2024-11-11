package exercise

import "fmt"

func add(a, b int, resultChan chan int) {
	resultChan <- a + b
}

func mul(a, b int, resultChan chan int) {
	resultChan <- a * b
}

func Exercise1() {
	var a, b int
	fmt.Print("첫 번째 정수를 입력하세요 : ")
	fmt.Scan(&a)
	fmt.Print("두 번째 정수를 입력하세요 : ")
	fmt.Scan(&b)

	addOutputChan := make(chan int)
	mulOutputChan := make(chan int)

	go add(a, b, addOutputChan)
	go mul(a, b, mulOutputChan)

	addresult := <-addOutputChan
	mulresult := <-mulOutputChan

	fmt.Println("덧셈 결과는 : ", addresult)
	fmt.Println("곱셈 결과는 : ", mulresult)
}

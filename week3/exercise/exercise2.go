package exercise

import "fmt"

func Exercise2() {
	var a, b int
	inputChan1 := make(chan int)
	inputChan2 := make(chan int)
	outputChan1 := make(chan int)
	outputChan2 := make(chan int)
	finishChan := make(chan int)

	go func(inputChan1, inputChan2, finishChan chan int) {
		for {
			select {
			case a := <-inputChan1:
				b := <-inputChan2
				outputChan1 <- a + b
				outputChan2 <- a * b
			case <-finishChan:
				return
			}
		}
	}(inputChan1, inputChan2, finishChan)
	fmt.Print("첫 번째 정수를 입력하세요 : ")
	fmt.Scan(&a)
	fmt.Print("두 번째 정수를 입력하세요 : ")
	fmt.Scan(&b)

	inputChan1 <- a
	inputChan2 <- b

	sum := <-outputChan1
	mul := <-outputChan2

	fmt.Println("덧셈 결과는 : ", sum)
	fmt.Println("곱셈 결과는 : ", mul)

	finishChan <- 1
}

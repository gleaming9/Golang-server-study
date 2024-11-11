package example

import "fmt"

func FuncCall() {
	inputChan := make(chan int, 10)
	outputChan := make(chan int, 10)
	exitChan := make(chan int, 1)
	go func(inputChan, exitChan chan int) { //선언부
		for {
			select {
			case x := <-inputChan:
				outputChan <- x * x
			case _ = <-exitChan:
				return
			}
		}
	}(inputChan, exitChan) //파라미터 전달

	for i := 0; i < 10; i++ {
		inputChan <- i
	}
	for i := 0; i < 10; i++ {
		fmt.Println(<-outputChan)
	}
	exitChan <- 1
}

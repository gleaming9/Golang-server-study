package example

import "fmt"

func squareIt2(inputChan, outputChan chan int) {
	for x := range inputChan {
		outputChan <- x * x
	}
}

func Channel() {
	inputChannel := make(chan int)
	outputChannel := make(chan int, 10)
	go squareIt2(inputChannel, outputChannel)
	for i := 0; i < 10; i++ {
		inputChannel <- i
	}
	for i := 0; i < 10; i++ {
		fmt.Println(<-outputChannel)
	}
	close(inputChannel)
}

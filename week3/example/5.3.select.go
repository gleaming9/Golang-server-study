package example

import "fmt"

func squarerCuber(sqInChan, sqOutChan, cuInChan, cuOutChan, exitChan chan int) {
	var squareX int
	var cubeX int
	for {
		select {
		case squareX = <-sqInChan:
			sqOutChan <- squareX * squareX
		case cubeX = <-cuInChan:
			cuOutChan <- cubeX * cubeX * cubeX
		case <-exitChan:
			return
		}
	}
}

func Select() {
	sqInChan := make(chan int, 10)
	sqOutChan := make(chan int, 10)
	cuInChan := make(chan int, 10)
	cuOutChan := make(chan int, 10)
	exitChan := make(chan int)
	go squarerCuber(sqInChan, sqOutChan, cuInChan, cuOutChan, exitChan)

	for i := 0; i < 10; i++ {
		sqInChan <- i
		cuInChan <- i
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("squarer says %d\n", <-sqOutChan)
		fmt.Printf("cuber says %d\n", <-cuOutChan)
	}
	exitChan <- 0
}

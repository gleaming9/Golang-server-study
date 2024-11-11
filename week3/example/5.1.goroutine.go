package example

import (
	"fmt"
	"time"
)

func squareIt(x int) {
	fmt.Println(x * x)
}

func Goroutine() {
	go squareIt(2)
	time.Sleep(1 * time.Millisecond)
}

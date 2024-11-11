package example

import (
	"fmt"
	"time"
)

func Capture() {
	for i := 0; i < 10; i++ {
		go func(nonCapturedI int) {
			time.Sleep(1 * time.Millisecond)
			fmt.Println(nonCapturedI)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)
}

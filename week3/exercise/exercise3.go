package exercise

/*
#cgo LDFLAGS: -L. -lfibonacci
int fibonacci(int);
*/
import "C"
import "fmt"

func Exercise3() {
	fmt.Println(C.fibonacci(C.int(10)))
}

package example

/*
#cgo LDFLAGS: -L. -lfactorial
int factorial(int);
*/
import "C"
import "fmt"

func ConnectSharedLibrary() {
	fmt.Println(C.factorial(C.int(5)))
}

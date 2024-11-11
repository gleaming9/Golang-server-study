package exercise

/*
#cgo LDFLAGS: -L. -lstring_concat
#include <stdlib.h>

char* string_concat(const char* str1, const char* str2);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func Exercise4() {
	a := C.CString("qwer")
	b := C.CString("asdf")

	result := C.GoString(C.string_concat(a, b))

	fmt.Println("Concatenated string :", result)

	C.free(unsafe.Pointer(a))
	C.free(unsafe.Pointer(b))
}

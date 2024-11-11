package example

/*
#include <stdio.h>
#include <stdlib.h>//C.free 함수를 사용하기 위해 필수로 추가
void printString(char *str){
	printf("%s\n",str);
}
*/
import "C"
import "unsafe"

func Cstring() {
	//Go 문자열을 C 문자열로 변환
	a := C.CString("This is from Golang")
	C.printString(a)
	//문자열에 필요한 메모리 사본이 만들어지므로 C 문자열을 사용할 때는
	//메모리를 해제해야 하고 필요 이상으로 오래 남아있지 않도록 한다.
	//메모리 누수 방지, Go 가비지 컬렉터는 내가 생성한 이 포인터를 처리할 수 없다.
	C.free(unsafe.Pointer(a))
}

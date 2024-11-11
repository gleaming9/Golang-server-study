package example

/*
#include <stdio.h>
char *getName(int idx){
	if (idx == 1)
		return "Tanmay Bakshi";
	if (idx == 2)
		return "Baheer Kamal";
	return "Invalid index";
}
*/
import "C"
import "fmt"

func GoString() {
	//이 경우 C에서 반환된 캐릭터 포인터를 해제할 필요가 없다.
	//해당 값이 애초에 힙에 할당되지 않고 스택에만 저장되었기 때문이다.
	cstr := C.getName(C.int(2))
	fmt.Println(C.GoString(cstr))
}

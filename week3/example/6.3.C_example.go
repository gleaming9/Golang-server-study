package example

/*
#include <stdio.h>
void printHelloWorld(){
	printf("Hello, World!\n");
}
*/
import "C"

func C_example() {
	C.printHelloWorld()
}

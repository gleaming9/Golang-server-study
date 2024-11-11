package errorExample

import (
	"fmt"
	"os"
)

func readFile(fileName string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", err))
	}
	defer file.Close()

	// 파일 처리
	fmt.Println("File opened successfully")
}

func PanicRecover2() {
	readFile("errorExample/example.txt")
	fmt.Println("Continuing execution")
}

package exercise

import (
	"fmt"
)

func sumArray(arr []int) int {
	var sum int
	for _, v := range arr {
		sum += v
	}
	return sum
}

func findMaxMin(arr []int) (int, int) {
	minValue := 100
	maxValue := -1
	for _, v := range arr {
		if v < minValue {
			minValue = v
		}
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue, minValue
}

func Exercise3() {
	defer fmt.Println("프로그램이 종료되었습니다.")
	arr := []int{3, 5, 1, 2, 0}

	fmt.Println("배열의 총합: ", sumArray(arr))
	maxv, minv := findMaxMin(arr)
	fmt.Println("최대값: ", maxv)
	fmt.Println("최소값: ", minv)

	switch {
	case len(arr) < 3:
		fmt.Println("배열의 길이가 짧습니다.")
	case len(arr) == 3:
		fmt.Println("배열의 길이가 적당합니다.")
	case len(arr) > 3:
		fmt.Println("배열의 길이가 깁니다.")
	}
}

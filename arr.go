package main

import "fmt"

func main() {
	/*
		var score [10]int // 建立一個大小為10的陣列
		score[0] = 90
		score[1] = 89
		arr1 := [3]int{1, 2, 3}
		arr2 := [...]int{1, 1, 2, 3, 4, 5}
		arr3 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
		fmt.Println(arr1)
		fmt.Println(arr2)
		fmt.Println(arr3)
		fmt.Println(score[0])
		fmt.Println(score)
	*/
	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ { //for loop
		fmt.Printf("%d\n", arr[i])
	}
	for index, element := range arr { //for range
		fmt.Printf("%d: %d\n", index, element)
	}
}

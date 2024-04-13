package main

import "fmt"

func main() {
	//fmt.Print(2)   //不換行
	//fmt.Println(2) //換行
	var x int
	var y int
	//fmt.Scanln(&x) //input
	//fmt.Scanln(&y)
	fmt.Scanln(&x, &y) // 3 5  
	fmt.Println(x * y) //&x取得變數指標
}

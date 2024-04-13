package main

import "fmt"

func main() {
	/*
		宣告slice的方法
		b := make([]int, 5, 10)
		Variable := make([]Type, Len, Cap)
		容量 cap ，最大能容納的長度
	*/
	a := make([]int, 10) //設定 len:10。現在長度10了，容量雖然沒給，但最大容納長度當然不可能小於10吧，所以就是10了
	fmt.Println(a, len(a), cap(a), len(a) == 0, a == nil)

	b := make([]int, 5, 10) //設定 len:5、cap:10
	fmt.Println(b, len(b), cap(b), len(b) == 0, b == nil)

	var c = []int{} //初始化slice
	fmt.Println(c, len(c), cap(c), len(c) == 0, c == nil)

	var d []int //尚未實體化，此時等於nil
	fmt.Println(d, len(d), cap(d), len(d) == 0, d == nil)

	e := []string{"youtube.com", "facebook.com"} //直接賦值
	fmt.Println(e, len(e), cap(e), len(e) == 0, e == nil)
	x := []int{1, 2, 3}
	x = append(x, 4, 5, 6)
	fmt.Println(x)

	//slice vs array:在於array建立時需要設定長度，設定後無法調整，但是slice可以動態調整，使用append
}

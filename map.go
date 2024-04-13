package main

import "fmt"

func main() {
	passwords := make(map[string]int) //map[keyType]valueType
	fmt.Println(passwords)            // map[]
	fmt.Println(len(passwords))       // 0
	pass := map[string]string{
		"account": "321654987",
		"passw":   "23748053",
	}
	fmt.Println(pass)
	passwords["account"] = 3
	passwords["passw"] = 456
	fmt.Println(passwords)
	fmt.Println(passwords["account"])
	fmt.Println(passwords["passw"])
	for key, value := range pass {
		fmt.Println(key, value)
	}
}

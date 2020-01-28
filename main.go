package main

import (
	"fmt"
	. "godemo/lesson1"
)



func main() {
	fmt.Println("\nHello", GetInfo())
	//switchFunc()
	//slice
	var arr = []int{1, 22, 33, 44, 54, 34, 13, 14}
	fmt.Println(Sum(arr))

	Closure(100)

	PrintProcess()

	//ReadStr()
}


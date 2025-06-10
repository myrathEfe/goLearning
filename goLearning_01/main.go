package main

import (
	"fmt"
)

func main() {
	//Fixed Length,Same Type
	intArr := [...]int32{0, 1, 2}

	intSlice := []int32{4, 5, 6}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Println("**********************")
	fmt.Println(intArr)

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

}

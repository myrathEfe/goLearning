package main

import "fmt"

// Arrays and Slices
func main() {
	//Fixed Length,Same Type
	fmt.Println("**********************")
	//[...]int32 yazımı, derleyiciye eleman sayısını otomatik bul demektir.
	intArr := [...]int32{2, 6, 10}
	fmt.Println("Array :\t", intArr)
	for i, v := range intArr {
		fmt.Printf("Index is : %v, Value is : %v\n", i, v)
	}
	fmt.Println("**********************")
	intSlice := []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity is %v\n", len(intSlice), cap(intSlice))
	fmt.Println("Slice :\t", intSlice)
	intSlice = append(intSlice, 7)
	fmt.Printf("After appending 7, the length is %v with capacity is %v\n", len(intSlice), cap(intSlice))
	fmt.Println("Slice :\t", intSlice)
	fmt.Println(&intSlice[0])
	fmt.Println(&intSlice[1])
	fmt.Println(&intSlice[2])
	fmt.Println(&intSlice[3])
	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println("Slice :\t", intSlice)
	fmt.Println("**********************")

}

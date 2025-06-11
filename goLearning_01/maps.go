package main

import "fmt"

// Maps
func main() {
	var myMap map[string]int32 = make(map[string]int32)
	fmt.Println(myMap)
	var myMap2 = map[string]int32{"Murat": 21, "Efe": 20, "Veli": 25, "Eren": 29, "Berk": 30}

	var age, ok = myMap2["Murat"]
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Println("Invalid Name")
	}
	fmt.Println("**********************")
	var age1, ok1 = myMap2["Deniz"]
	if ok1 {
		fmt.Printf("The age is %v\n", age1)
	} else {
		fmt.Println("Invalid Name")
	}

	for name, age := range myMap2 {
		//Çıktı farklı sıralarda olabilir.
		fmt.Printf("The name is %v Age is :%v \n", name, age)
	}

}

package main

import "fmt"

type owner struct {
	name string
}

type gasEngine struct {
	kmpl  uint8
	litre uint8
	owner //embedded struct
}
type electricEngine struct {
	kmpkwh uint8
	kwh    uint8
	owner
}

/*func (receiver Type) MethodName() ReturnType {
	// method body
}
*/

func (e gasEngine) kilometersLeft() uint8 {
	return e.kmpl * e.litre
}

func (e electricEngine) kilometersLeft() uint8 {
	return e.kwh * e.kmpkwh
}

type engine interface {
	kilometersLeft() uint8
}

func canGo(e engine, kilometers uint8) {
	if kilometers > e.kilometersLeft() {
		fmt.Printf("No, the car can't go %v kilometers. It can only go %v kilometers.\n", kilometers, e.kilometersLeft())
	} else {
		fmt.Printf("Yes, the car can go %v kilometers. %v kilometers of range will remain.\n", kilometers, e.kilometersLeft()-kilometers)
	}
}

func main() {
	var myEngine = gasEngine{15, 20, owner{"Efe"}}
	var myEngine2 = electricEngine{20, 10, owner{"Efe"}}
	fmt.Println(myEngine.kmpl, myEngine.litre, myEngine.name)
	fmt.Println(myEngine2.kmpkwh, myEngine2.kwh, myEngine2.owner)
	fmt.Printf("Total kilometers left : %v\n", myEngine.kilometersLeft())
	fmt.Printf("Total kilometers left : %v\n", myEngine2.kilometersLeft())
	canGo(myEngine, 100)
	canGo(myEngine, 160)
	canGo(myEngine2, 180)
	canGo(myEngine2, 210)
}

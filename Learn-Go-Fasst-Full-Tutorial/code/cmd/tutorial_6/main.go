package main

import "fmt"

type gasEngin struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

type owner struct {
	name string
}

func (e gasEngin) milesLeft() uint8 {
	return e.mpg * e.gallons
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

type engine interface {
	milesLeft() uint8
}

// Interface

func main() {
	//var myEngine gasEngin = gasEngin{25, 15, owner{"Alex"}}
	//fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name)
	//fmt.Printf("Total miles left in Tank: %d\n", myEngine.milesLeft())
	//fmt.Printf("Total miles left in Battery: %d\n", myEngine.milesLeft())
	var myEngine electricEngine = electricEngine{25, 24}
	canMakeIt(myEngine, 50)
}

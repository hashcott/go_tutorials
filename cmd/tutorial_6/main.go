package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.gallons
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if e.milesLeft() >= miles {
		fmt.Println("We can make it!")
	} else {
		fmt.Println("Need to fill up!")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{25, 18}
	canMakeIt(myEngine, 50)
	var myEngine2 electricEngine = electricEngine{25, 18}
	canMakeIt(myEngine2, 200)
}

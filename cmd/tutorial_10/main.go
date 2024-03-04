package main

import "fmt"

type gasEngine struct {
	gallons float32
	mpg     float32
}

type electricEngine struct {
	kwh   float32
	mpkwh float32
}

type car[X gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   X
}

func isEmpty[X any](slice []X) bool {
	return len(slice) == 0
}

func main() {
	var gasCar = car[gasEngine]{
		carMake:  "Toyota",
		carModel: "Corolla",
		engine: gasEngine{
			gallons: 10,
			mpg:     30,
		},
	}
	var electricCar = car[electricEngine]{
		carMake:  "Tesla",
		carModel: "Model 3",
		engine: electricEngine{
			kwh:   60,
			mpkwh: 4,
		},
	}

	fmt.Println(gasCar, electricCar)

	var float32Slice = []float32{1.1, 2.2, 3.3}
	fmt.Println(isEmpty(float32Slice))
}

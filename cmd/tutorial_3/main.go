package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	var printValue string = "Hello, World!"
	printMe(printValue)

	var numerator int = 10
	var denominator int = 2
	if result, remainder, err := intDivision(numerator, denominator); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%d divided by %d equals %d with a remainder of %d\n", numerator, denominator, result, remainder)
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("division by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}

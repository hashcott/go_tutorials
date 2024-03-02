package main

import "fmt"

func main() {
	var p *int = new(int)
	var i int
	fmt.Printf("The value p pointer to is: %v", *p)
	fmt.Printf("\nThe value if i is: %v", i)
	p = &i
	*p = 1
	fmt.Printf("\nThe value p pointer to is: %v", *p)
	fmt.Printf("\nThe value if i is: %v", i)

	// slice
	var slice = []int{1, 2, 3, 4, 5}
	var sliceCopy = slice
	sliceCopy[0] = 100
	fmt.Printf("\nSlice: %v", slice)
	fmt.Printf("\nSliceCopy: %v", sliceCopy)

	// argument passing
	var thing1 = [5]int{1, 2, 3, 4, 5}
	var thing2 = square(&thing1)
	fmt.Printf("\nThing1: %v", thing1)
	fmt.Printf("\nThing2: %v", thing2)
}

func square(thing2 *[5]int) [5]int {
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}

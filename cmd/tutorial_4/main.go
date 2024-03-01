package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		[]Arrays
		Fixed length
		Same type
		Indexable
		Contiguous memory
	*/
	var intArr [3]int
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	intArr2 := [...]int{1, 2, 3}
	fmt.Println(intArr2)

	var intSlice []int = []int{4, 5, 6}
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice), cap(intSlice))
	// fmt.Println(intSlice[4])

	var intSlice2 []int = []int{3, 5}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var myMap map[string]uint = make(map[string]uint)
	fmt.Println(myMap)

	var myMap2 = map[string]uint{"Harry": 23, "Ron": 22, "Hermione": 21}
	fmt.Println(myMap2["Harry"])
	var age, ok = myMap2["Ginny"]
	fmt.Println(age, ok)

	for name, age := range myMap2 {
		fmt.Println(name, age)
	}

	for i, v := range intArr {
		fmt.Println(i, v)
	}

	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var n int = 10000000
	fmt.Printf("Total time without preallocation: %v\n", timeLoop([]int{}, n))
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(make([]int, 0, n), n))
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}

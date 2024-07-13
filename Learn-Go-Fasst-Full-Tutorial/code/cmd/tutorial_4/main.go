package main

import (
	"fmt"
	"time"
)

/*
	[]Arrays
		- Fixed Length
		- Same Type
		- Indexable
		- Contiguous in Memory
*/

/*
	Slices wrap arrays to give a more general,
	powerful, and convenient interface to sequences of data.
*/

/*
	Map are key value pairs.
*/

func main() {

	// Array
	var intArr [3]int32
	intArr[1] = 123

	fmt.Println(intArr[0])
	fmt.Println(intArr[2:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	floatArr := [3]float32{1.2, 3.4, 4.5}
	fmt.Println(floatArr)

	boolArr := [...]bool{true, false, true}
	fmt.Println(boolArr)

	// Slice
	intSlice := []int32{7, 8, 9}
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with cap is %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 123)
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with cap is %v\n", len(intSlice), cap(intSlice))

	floatSlice := make([]float32, 3, 8)
	fmt.Println(floatSlice)
	fmt.Printf("The length is %v with cap is %v\n", len(floatSlice), cap(floatSlice))

	// map
	myMap := make(map[string]int)
	fmt.Println(myMap)
	myMap2 := map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Adam"])
	age, ok := myMap2["Sarah"]
	if ok {
		fmt.Println(age)
	} else {
		fmt.Println("not found key")
	}
	// delete(myMap2, "Sarah")
	// fmt.Println(myMap2)

	for name, age := range myMap2 {
		fmt.Println(name, age)
	}

	for i, v := range intArr {
		fmt.Println(i, v)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// test time
	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v \n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v \n", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}

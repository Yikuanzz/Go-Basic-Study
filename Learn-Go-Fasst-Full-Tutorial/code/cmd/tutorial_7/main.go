package main

import "fmt"

func main() {

	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value of p is %v\n", *p)
	fmt.Printf("The value if i is %v\n", i)

	p = &i
	*p = 1
	fmt.Printf("The value of p is %v\n", *p)
	fmt.Printf("The value if i is %v\n", i)

	var slice = []int32{1, 2, 3}
	var sliceCopy = slice // Slice just copy the address
	sliceCopy[2] = 4
	fmt.Printf("The value of slice is %v\n", slice)
	fmt.Printf("The value of sliceCopy is %v\n", sliceCopy)

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("The memory location of the thing1 array is %p\n", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("The result is %v\n", result)
	fmt.Printf("The value of thing1 is %v\n", thing1)
}

// There some different between value parameter and pointer parameter.
func square(thing2 *[5]float64) [5]float64 {
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}

## Introduction to Golang

* Statically Typed Language.

* Strongly Typed Language.
* GO is Complied.
* Fast Compile Time.
* Built in Concurrency.
* Simplicity.



We need to know the concepts of modules and packages. 

* Package is just a folder that contains a collection of go files.
* Module is to collect these packages.

Than I download go in Linux system and use the command to initialize the go mod: `go mod init go_tutorials`.

```go
package main

import "fmt"

func main(){
	fmt.Println("Hello World!")
}
```



## Constant Variable  and Basic Data Types

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Need to use
	var intNum int64
	intNum += 1

	var floatNum float64
	floatNum = 3.14

	var myString string
	myString = "Yikuanzz"

	var myBool bool
	myBool = true

	fmt.Printf("int number: %d\n", intNum)
	fmt.Printf("float number: %.2f\n", floatNum)
	fmt.Printf("my string: %s\n", myString)
	fmt.Println(utf8.RuneCountInString("Y"))
	fmt.Printf("myBool: %t\n", myBool)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const myConst string = "const value"
	const pi float64 = 3.1
	fmt.Println(myConst, pi)
}
```



## Functions and Control Structures

```go
package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe("Hello World!")

	numerator, denominator := 11, 3
	result, remainder, err := intDivision(numerator, denominator)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("The result of the interger division is %v with remainder %v.", result, remainder)
	}

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return numerator, denominator, err
	}

	var result = numerator / denominator
	var remainder = numerator % denominator
	return result, remainder, nil
}

```



## Arrays, Slices, Maps and Loops

```go
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

```



## Strings, Runes and Bytes

- `%v` represents the named value in its default format.
- `%d` expects the named value to be an integer type.
- `%f` expects the named value to be a float type.
- `%T` represents the type for the named value.

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "Yikuanzz"
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Printf("\n The length of String is %d\n", len(myString))

	var strBuilder strings.Builder
	var strSlice = []string{"s", "u", "b", "s", "c"}
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)

}

```



## Structs and Interfaces

```go
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

```



## Pointers

```go
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

```



## Goroutines

We need to know about concurrency and parallel execution.

* short answer: Concurrency is two lines of customers ordering from a single cashier (lines take turns ordering); Parallelism is two lines of customers ordering from two cashiers (each line gets its own cashier) .



**Mutex (Mutual Exclusion)**:

- Also known as an exclusive lock, a Mutex ensures that only one goroutine can execute a critical section of code at a time.
- It provides two main methods: `Lock()` (to acquire the lock) and `Unlock()` (to release the lock).
- The code between `Lock()` and `Unlock()` is called the **critical section**, and it is thread-safe. Only one goroutine can execute this section at any given time.



**RWMutex (Read-Write Mutex)**:

- RWMutex allows multiple readers to access a resource simultaneously or a single writer to have exclusive access.
- It has two types of locks:
  - `RLock()` and `RUnlock()`: Used for read operations. Multiple readers can hold the read lock concurrently.
  - `Lock()` and `Unlock()`: Used for write operations. Only one writer can hold the write lock at a time.
- RWMutex is useful when there are many concurrent readers and fewer writers.



```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

var wg = sync.WaitGroup{} // WaitGroup just counters.
// var m = sync.Mutex{}    // Use Lock method and Unlock method to protect the code.
var m = sync.RWMutex{} // Use RLock method and RUnlock method to protect the code.
var results []string

func main() {
	t0 := time.Now()

	for i := 0; i < len(dbData); i++ {
		/*
			When we spawn a new go routine we
			make sure to add one to the counter.
			And we will call the done method
			at the end inside the function we used.
		*/
		wg.Add(1)
		go dbCall(i)
	}
	/*
		We call the weight method this method is going to
		wait for the counter to go back down to zero.
	*/
	wg.Wait()
	fmt.Printf("Total execution time: %v\n", time.Since(t0))
	fmt.Printf("The results: %v\n", results)
}

func dbCall(i int) {
	/*
		Actually we need to mutex to control the writing to our slice.
		We should care about corrupting memory.

		But there is still a question that the mutex may destroy the concurrency of database calls.
		To solve this problem, we can use RWMutex.
	*/
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	//fmt.Println("The result from the database is ", dbData[i])
	//m.Lock() // It'll wait for another go routine release the lock.
	//results = append(results, dbData[i])
	//m.Unlock()

	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("The current results are: %v\n", results)
	m.RUnlock()
}

```



## Channels

In Go, a **channel** is a core type that acts as a communication mechanism between concurrent goroutines. Think of it as a pipeline for exchanging data. Here are some key points about channels:

1. **Declaration and Initialization**:
   - You create a channel using the `make` function: `ch := make(chan int)`.
   - The type of data that the channel can carry is specified (e.g., `int` in the example above).
2. **Sending and Receiving Data**:
   - To send data into a channel, use the arrow operator (`<-`): `ch <- value`.
   - To receive data from a channel, use the same arrow operator: `value := <-ch`.
3. **Blocking Behavior**:
   - Sending or receiving from a channel blocks until the other side is ready.
   - This synchronization ensures safe communication between goroutines.
4. **Closing a Channel**:
   - You can close a channel using `close(ch)`.
   - Closed channels can still be read from, but any subsequent writes will panic.
5. **Buffered Channels**:
   - By adding a buffer size when creating a channel (e.g., `ch := make(chan int, 10)`), you create a buffered channel.
   - Buffered channels allow multiple goroutines to send or receive without immediate blocking.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		If we run this code we will get the deadlock!
		When we write to an unbuffered channel the code will
		block until something else reads from it.
		So the channel will be waiting forever.
	*/
	//var c = make(chan int)
	//c <- 1      // put 1 into the channel and it is [1].
	//var i = <-c // pop 1 from the channel and it is [].
	//fmt.Println(i)

	var c = make(chan int)
	go process(c)
	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}

func process(c chan int) {
	/*
		We have to close the channel.
		Otherwise, there will be deadlock because
		reader will keep reading and wait for the next message.
		But nobody will send messages to it.
	*/
	defer close(c)
	for i := 0; i < 10; i++ {
		c <- i
	}
	fmt.Println("Exiting process")
}

```



## Generics

**Generics** in Go allow you to write code that’s independent of specific types. With generics, you can declare functions or types that work with any set of types provided by calling code. Here’s a brief overview:

1. **Type Parameters**:
   - You can define functions or types with type parameters (placeholders for actual types).
   - Example: `func SumT Numeric T { ... }`
2. **Type Constraints**:
   - You can specify constraints on type parameters (e.g., `Numeric`, `Comparable`, etc.).
   - Example: `func SumT Numeric T { ... }`
3. **Interface Types**:
   - Generics allow defining interface types as sets of types (including those without methods).
   - Example: `type Container[T any] interface { ... }`

```go
package main

import (
	"fmt"
)

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice[int](intSlice))
	fmt.Println(isEmpty(intSlice))

	var float32Slice = []float32{1.1, 2.2, 3.3}
	fmt.Println(sumSlice[float32](float32Slice))
	fmt.Println(isEmpty(float32Slice))
}

// Or we can use any type in here.
func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

```



Also we can use generics in this example.

```go
package main

import "fmt"

type gasEngine struct{
    gallons float32
    mpg float32
}

type electricEngine struct{
    kwh float32
    mpkwh float32
}

type car[T gasEngine | electricEngine] stuct{
    carName string
    carModel string
    engine T
}

func main(){
    var gasCar = car[gasEngine]{
        carName : "Honda",
        carModel : "Civic",
        engine: gasEngine{
            gallons: 12.4,
            mpg: 40,
        },
    }
    
    var electricCar = car[electricEngine]{
        carName: "Tesla",
        carModel: "Model 3",
        engine: electricEngine{
            kwh: 57.3,
            mpkwh: 4,17,
        },
    }
}
```












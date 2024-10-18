package main

import (
	"fmt"
)

func main() {

	var ar1 = [3]int{5, 9, 10}
	ar2 := [2]string{"sanjiv", "kumar"}

	var ar3 = [...]int{5, 9, 10, 3, 19, 32}

	fmt.Println("Values of ar1 = ", ar1)
	fmt.Println("Values of ar2 = ", ar2)
	fmt.Println("Values of ar2 = ", ar3)

	// ----------------------------------------------

	arr1 := [5]int{}              //not initialized, default integers will be zero
	arr2 := [5]int{1, 2}          //partially initialized
	arr3 := [5]int{1, 2, 3, 4, 5} //fully initialized

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3[3])

	// ----------------------------------------------

	// initialize specific element of the array

	ar4 := [5]int{2: 40, 4: 90} // ar[2]=40, ar[4]=90, remaining will be zero

	fmt.Println("Values of ar4 = ", ar4)

	// finding length of an array

	fmt.Println("Length of ar4 = ", len(ar4))

	// creating an empty array
	empArr := [...]int{}

	fmt.Println(empArr, "Length = ", len(empArr))

	// inserting elements in the empty array of length=0
	// -- it gives a runtime error: index out of range [0] with length 0
	for i := 0; i < 10; i++ {
		empArr[i] = i
	}
	fmt.Println(empArr, "Length = ", len(empArr))

}

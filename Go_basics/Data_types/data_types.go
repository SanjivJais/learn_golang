package main

import "fmt"

func integer_type() {
	// signed and unsigned integers

	// signed integers using "int"
	var x int = 56
	var y int = -9

	fmt.Printf("Value of x is %d and its type is %T\nValue of y is %v and its type is %T", x, x, y, y)

	// unsigned integers using "uint"
	var z uint = 40 //can only store non-negative numbers
	fmt.Printf("\nValue of z = %v, Type: %T", z, z)

	/*

		Default value of int is 0

		✅ Go has five keywords/types of signed/unsigned integers:
		1. int (depending on the system, it could be 32bit or 62bit ranging from -2147483648 to 2147483647 in 32 bit systems and
		-9223372036854775808 to 9223372036854775807 in 64 bit systems)

		2. int8 (8bit integer from -128 to 127)

		Similarly:
		3. int 16
		4. int 32
		5. int 64


	*/
}




func float_type() {
	// float type has two keywords: float32, float64

	var x float32 = -3.1415
	fmt.Printf("\nValue of x = %v, Type: %T", x, x)

	/*

		1) float32 ranges from -3.4e+38 to 3.4e+38
		2) float64 ranges from -1.7e+308 to +1.7e+308.

	*/

}



func string_type(){

	// The string data type is used to store a sequence of characters (text). 
	// String values must be surrounded by double quotes:

	var name string
	name = "Sanjiv"

	// OR, name:="Sanjiv"

	fmt.Printf("\nMy name is %v, data type is %T", name, name)

}




func bool_type(){
	// A boolean data type is declared with the bool keyword and can only take the values true or false.
	// Default value of a bool variable is `false` 

	var b1 bool
	b2:= true

	fmt.Printf("\nValue of b1: %v, Type: %T\nValue of b2: %v, Type: %T", b1, b1, b2, b2)
	
}






func main() {
	integer_type()
	float_type()
	string_type()
	bool_type()
}

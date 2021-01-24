package main

import "fmt"

func main() {
	var a = 4
	var b = 5.2

	/*
		a = b  ===>  Compile-time ERROR: cannot use b (type float64) as type int in assignment
		Type of a variable CANNOT be changed once declared
	*/

	a = int(b) // Converts type FLOAT64 to INT for reassignment

	fmt.Println(a, b) // 5 5.2

	var x int
	// x = "5"  ===> ERROR: cannot use "5" (type untyped string) as type int in assignment

	// x = int("5")  ===> ERROR: cannot convert "5" (type untyped string) to type int
	_ = x // (_) === Blank Indentifier to silence unused variable error

	var value int
	var price float64
	var name string
	var done bool
	fmt.Println(value, price, name, done) // 0 0  false
}

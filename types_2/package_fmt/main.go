package main

import "fmt"

func main() {

	/*
		FOR fmt.Printf("... \n")

		%d		= INTERGER
		%f		= FLOAT
		%.3f	= FLOAT with 3 DECIMAL PLACES
		%s		= STRING
		%q		= QUOTED STRING ("")
		%v		= UNIVERSAL VERB (applies to all types)
		%#v		= ARRAY and SLICE type with VALUES
		%T		= TYPE
	*/

	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}

	// 1.
	fmt.Printf("x is %d, y is %f, z is %s\n", x, y, z) // x is 10, y is 15.500000, z is Gophers
	fmt.Printf("score is %#v\n", score)                // score is []int{10, 20, 30}

	// 2.
	fmt.Printf("z is %q\n", z) // z is "Gophers"

	// 3.
	fmt.Printf("x is %v, y is %v, z is %v, score is %v\n", x, y, z, score) // x is 10, y is 15.5, z is Gophers, score is [10 20 30]

	// 4.
	fmt.Printf("y type: %T, score type: %T\n", y, score) // y type: float64, score type: []int

	const X float64 = 1.422349587101

	fmt.Printf("This is const X with 4 decimal points = %.4f \n", X) // This is const X with 4 decimal points = 1.4223

	shape := "circle"
	radius := 3.2
	const pi float64 = 3.14159
	circumference := 2 * pi * radius

	arr := [...]int{2, 3, 4, 5} // %#v ==> [4]int{2, 3, 4, 5}
	_ = arr

	fmt.Printf("Shape: %q\n", shape)                                                    // Shape: "circle"
	fmt.Printf("Circle's circumference with radius %f is %#v\n", radius, circumference) // Circle's circumference with radius 3.200000 is 20.106176
	_ = shape

}

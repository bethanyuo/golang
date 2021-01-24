package main

import "fmt"

func main() {
	// var a int
	// var b float64
	// var c bool
	// var d string

	var (
		a int
		b float64
		c bool
		d string
	)

	a = 12 // Reassignment with = sign
	a = 2

	// x := 20
	// y := 15.5
	// z := "Gopher!"

	x, y, z := 20, 15.5, "Gopher!"

	x = 30 // Reassignment with = sign

	fmt.Println(a, b, c, d, x, y, z) // 2 0 false  30 15.5 Gopher!
	/*
		_, _, _, _, _, _, _ = a, b, c, d, x, y, z //using blank identifier to mute the unused variable error
		// _ stays always on the left side of =
	*/

	var a1 float64 = 7.1

	x1, y1 := true, 3.7

	// a1, x1 := 5.5, false // no new variables on left side of :=

	a1, x1 = 5.5, false // Must REMOVE (:) from = as we are not redeclaring but reassigning (=)

	_, _, _ = a1, x1, y1

	// name := 'Golang'  // more than one character in rune literal

	name := "Golang" // A string is initialized only using double quotes ("")

	fmt.Println(name) // Golang
}

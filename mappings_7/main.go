/////////////////////////////////
// Maps in Go
// Go Playground: https://play.golang.org/p/BMtPVKuOwEQ
/////////////////////////////////

package main

import "fmt"

func main() {
	// var m1 map[nil]string
	// //the zero value of a map is nil

	// // the zero value of a map is nil
	// fmt.Printf("%#v\n", m1) // -> map[string]string(nil).

	// 1.
	var m1 map[float64]bool
	fmt.Printf("m1 type: %T, m1 value: %#v\n", m1, m1) // m1 type: map[float64]bool, m1 value: map[float64]bool(nil)

	// 2.
	m2 := map[int]string{1: "Sting", 2: "Queen"} // Initialize the map using a map literal with two key:value pairs

	// 3.
	m2[10] = "Abba"

	// 4.
	fmt.Println(m2[2])   // existing key  ==>  Queen
	fmt.Println(m2[10])  // existing key  ==>  Abba
	fmt.Println(m2[100]) // non-existing key  ==>  EMPTY String / Space

	var m12 map[int]string
	_ = m12

	var m map[int]bool
	// m[5] = true // ERROR: panic: assignment to entry in nil map

	mm := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}

	// fmt.Println(mm == m3) // ERROR: invalid operation: mm == m3 (map can only be compared to nil)

	_, _, _ = m, mm, m3

	// 1.
	_m := map[int]bool{1: true, 2: false, 3: false}

	// 2.
	delete(_m, 2) // Delete key:value pair

	// 3.
	for k, v := range _m {
		fmt.Printf("k: %d, v: %t\n", k, v) // Iterate over the mapping
	}
	// Result:
	/*
		k: 1, v: true
		k: 3, v: false
	*/
}

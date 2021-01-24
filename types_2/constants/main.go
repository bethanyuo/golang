package main

import (
	"fmt"
	"math"
)

func main() {

	// const daysPerWeek int = 7
	// const lightSpeed int = 299792458
	// const pi float64 = 3.14159

	const (
		daysPerWeek = 7
		lightSpeed  = 299792458
		pi          = 3.14159
	)

	const secPerDay = 3600 * 24
	const daysYear = 365

	fmt.Printf("There are %d seconds in a year.\n", secPerDay*daysYear) // There are 31536000 seconds in a year

	// _, _, _ = daysPerWeek, lightSpeed, pi // Unnecessary for CONSTANTS

	// const x int = 10

	// declaring a constant of type slice int ([]int)
	// const m = []int{1: 3, 4: 5, 6: 8} // const initializer []int literal is not a constant

	var m = []int{1: 3, 4: 5, 6: 8}
	_ = m

	///////////////////////////////////////////////
	/*
		const a int = 7
		const b float64 = 5.6
		const c = a * b     	 	// invalid operation: a * b (mismatched types int and float64)
	*/

	const a int = 7 // OR, make const a UNTYPED ==> const a = 7  ===> const c = a * b
	const b float64 = 5.6
	const c = float64(a) * b

	/*
		x := 8						 // ERRROR ->  You cannot initiate a constant at runtime (constants belong to compile-time)
		const xc int = x			 // const initializer x is not a constant
	*/

	// variables belong to runtime
	x := 8
	var xc int = x
	_ = xc

	//	const noIPv6 = math.Pow(2, 128)   // ERROR: const initializer math.Pow(2, 128) is not a constant
	// ERRROR ->  You cannot initiate a constant at runtime (constants belong to compile-time)

	// function calls belong to runtime
	var noIPv6 = math.Pow(2, 128)
	_ = noIPv6

	const (
		//iota starts from zero
		Jun = iota + 6
		Jul //automatically incremented by one
		Aug
	)

	fmt.Println(Jun, Jul, Aug) // 6 7 8
}

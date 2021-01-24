package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 3
	var f float64 = 3.2

	i, f = int(f), float64(i)

	var a = 3
	var b = 3.2
	var s1, s2 = "3.14", "5"

	// 1. int to string
	s := strconv.Itoa(a)
	fmt.Printf("s Type is %T, s value is %q\n", s, s) // s Type is string, s value is "3"

	// 2. string to int
	is, err := strconv.Atoi(s2)
	if err == nil {
		fmt.Printf("i type is %T, i value is %v\n", is, is) // i type is int, i value is 5
	} else {
		fmt.Println("Cannot convert string to int.")
	}

	// 3. float64 to string
	ss1 := fmt.Sprintf("%f", b)
	fmt.Printf("ss1's type: %T, ss1's value: %s\n", ss1, ss1) // ss1's type: string, ss1's value: 3.200000

	// 4. string to float64
	f1, err1 := strconv.ParseFloat(s1, 64)
	if err1 == nil {
		fmt.Printf("f1's type: %T, f1's value: %v\n", f1, f1) // f1's type: float64, f1's value: 3.14
	} else {
		fmt.Println("Cannot convert string to float64.")
	}

	x, y := 4, 5.1

	z := x * int(y)
	fmt.Println(z) // 20  OR  20.4 if using float64()

	c := float64(5)
	d := 6.2 * c
	fmt.Println(d) // 31

	// _, _ = a, b

	distance := 149.6 * 1000_000_000
	speed := 299_792_458

	time := distance / float64(speed)
	_, _ = distance, speed

	//  time = distance / speed

	fmt.Printf("The distance from the Sun to Earth is %.0fm, and the speed of light is %dm/s.\n", distance, speed)
	// The distance from the Sun to Earth is 149600000000m, and the speed of light is 299792458m/s.

	fmt.Printf("The Sunlight reaches Earth in %.0f seconds.\n", time) // The Sunlight reaches Earth in 499 seconds.

	x1, y1 := 0.1, 5
	var z1 float64

	// Write the correct logical operator (&&, || , !)
	// inside the expression so that result1 will be false and result2 will be true.

	result1 := x1 < z1 || int(x1) != int(z1)
	fmt.Println(result1) // false

	fmt.Println(int(x)) // 0
	fmt.Println(int(z)) // 0

	result2 := y1 == 1*5 && int(z1) == 0 // OR ||
	fmt.Println(result2)                 // true
}

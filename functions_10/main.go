package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func cube(perimeter float64) float64 {
	c := math.Pow(perimeter, 3)
	fmt.Println(c)
	return c
}

func f1(n uint) int { // Simply Summation of all INTs upto N

	sum := 0
	switch {
	case n == 0 || n == 1:
		sum += int(n)
	default:
		for i := 1; i <= int(n); i++ {
			//fact *= i
			sum += i
		}

	}
	return sum
}

func f2(n uint) (int, int) { // Factorial and Sum

	fact := 1
	sum := 0

	for i := 1; i <= int(n); i++ {
		fact *= i
		sum += i
	}
	return fact, sum
}

func myFunc(n string) int {
	// convert string to int
	si, err := strconv.Atoi(n)
	sii, _ := strconv.Atoi(n + n)
	siii, _ := strconv.Atoi(n + n + n)

	// error handling
	if err != nil {
		fmt.Println("Can not convert string to int.")
		os.Exit(1) //exit the program
	}
	sum := si + sii + siii
	fmt.Printf("The value of (n + nn + nn) is %v\n", sum)
	return sum
}

func sum1(a ...int) (sum int) { // Variadic Func ==> Sum of all values of type int
	// sum = 0
	for _, v := range a {
		sum += v
	}
	return
}

func print(msg string) {
	fmt.Println(msg)
}

func integer(v int) {
	fmt.Println(v)
}

func searchItem(mySlice []string, myStr string) bool { // Case-sensitive Search
	for _, v := range mySlice {
		if v == myStr {
			return true
		}

	}
	return false
}

func searchItem2(mySlice []string, myStr string) bool { // Case-insensitive Search
	for _, v := range mySlice {
		if strings.EqualFold(v, myStr) {
			return true
		}

	}
	return false
}

func main() {

	cube(5.4) // 157.46400000000003

	sum := f1(5)
	fmt.Println(sum) // 15

	f, s := f2(4)
	fmt.Println("f:", f, "s:", s) // f: 24 s: 10

	f, s = f2(10)
	fmt.Println("f:", f, "s:", s) // f: 3628800 s: 55

	myFunc("2") // The value of (n + nn + nn) is 246

	s1 := sum1(1, 2, 30)
	fmt.Println(s1) // 33

	i := integer                     // i == the function integer ==> Do not use () at the end
	fmt.Printf("Func Type: %T\n", i) // Func Type: func(int)
	integer(8)                       // 8

	animals := []string{"lion", "tiger", "bear"}

	result := searchItem(animals, "bear")
	fmt.Println(result) // => true

	result = searchItem(animals, "BEAR")
	fmt.Println(result) // => false

	result = searchItem(animals, "pig")
	fmt.Println(result) // => false

	result = searchItem2(animals, "BEAR")
	fmt.Println(result) // => true

	result = searchItem2(animals, "lION")
	fmt.Println(result) // => true

	result = searchItem2(animals, "pig")
	fmt.Println(result) // => false

	defer print("The Go gopher is the iconic mascot of the Go project.")
	fmt.Println("Hello, Go playground!")
	// Result - Deferrence
	/*
		Hello, Go playground!
		The Go gopher is the iconic mascot of the Go project.
	*/
}

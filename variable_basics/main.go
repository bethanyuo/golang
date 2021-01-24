package main

import "fmt"

func main() {
	var age int = 30 // Do not need int declaration
	fmt.Println("Your age:", age)

	var name = "Dan" // OR could use ===> var name string = "Dan"
	// name := "Sam"  ===> ERROR: no new variables on left side of :=
	// Can ony reassign variable name with = sign (name = "Sam")

	fmt.Println("Your name:", name)

	s := "Learning Golang!" // new variable assignment
	fmt.Println(s)

	// can also use ===> _ = variableName ===> to silence an unused variable and avoid runtime error

	car, cost := "Mercedes", 30000
	fmt.Println(car, cost)    // Mercedes 30000
	car, year := "Audi", 2020 // car variable can be reassigned with := as long as another new variable is present (year)
	_ = year
	fmt.Println(car, cost) // Audi 30000

	var opened = false
	opened, file := true, "a.txt"

	_, _ = opened, file // unused variables (opened, file) are now silenced

	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender) // 0  false ===> ZERO for empty (numerical) float64, with SPACE afterwards for the empty string, and FALSE for the unassignned boolean

	var a, b, c int
	fmt.Println(a, b, c) // 0 0 0 ===> all ZEROES for the unassignned int variables

	var i, j int
	i, j = 5, 8

	j, i = i, j

	fmt.Println(i, j) // 8 5 ===> Switched variable values

	sum := 5 + 2.3
	fmt.Println(sum) // 7.3

}

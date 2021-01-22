package main

import "fmt"

const secondsInHour = 3600 // variables CANNOT use underscores =/=> seconds_in_hour (NOT ALLOWED)

func main() {
	fmt.Println("Hello Go World")
	distance := 60.8

	fmt.Printf("The distance in miles is %f \n", distance*0.62137) // \n = works as a line break within terminal
	// between code and terminal base
}

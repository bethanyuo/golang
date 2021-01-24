package main

import "fmt"

type duration int
type mile float64
type kilometer float64

const m2km = 1.609 //km ==> 1mil == 1.609km  // FLOAT64

func main() {
	var hour duration

	// print out the value of the variable hour
	//print out the type of the variable hour

	fmt.Printf("hour's type: %T, hour's value: %v\n", hour, hour) // our's type: main.duration, hour's value: 0

	//assign 3600 to the variable hour using the = operator
	hour = 3600

	//print out the value of hour
	fmt.Printf("hour's value: %v\n", hour) // hour's value: 3600

	var oneHour duration = 3600
	minute := 60

	// Convert minute which is of type int() to type duration()
	fmt.Println(oneHour != duration(minute)) // true

	// m2km := 1.609 //km ==> 1mil == 1.609km

	var mileBerlinToParis mile = 655.3 //mil

	var kmBerlinToParis kilometer

	kmBerlinToParis = kilometer(m2km * mileBerlinToParis) // Declared CONSTANTS used removes TYPE error msg

	fmt.Printf("The distance in km between Berlin and Paris is %.2fkm.\n", kmBerlinToParis)
	// The distance in km between Berlin and Paris is 1054.38km.

}

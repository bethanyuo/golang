package main

import "fmt"

// 1.
type person struct {
	name   string
	age    int
	colors []string
	score  grades // Embedded struct grades
}

type grades struct {
	grade  int
	course string
}

func main() {
	// 2.
	me := person{name: "Marius", age: 30, colors: []string{"red", "yellow"}, score: grades{grade: 88, course: "Math"}}
	you := person{
		name:   "Maria",
		age:    22,
		score:  grades{grade: 75, course: "French"},
		colors: []string{"pink", "blue"},
	}

	// 3.
	fmt.Printf("%v\n", me)   // {Marius 30 [red yellow]}
	fmt.Printf("%#v\n", me)  // main.person{name:"Marius", age:30, colors:[]string{"red", "yellow"}}
	fmt.Printf("%+v\n", you) // {name:Maria age:22 colors:[pink blue]}

	// 1.
	me.name = "Andrei"

	// 2.
	var colors []string = you.colors
	fmt.Printf("Type: %T, Value: %v\n", colors, colors) // Type: []string, Value: [pink blue]

	// 3.
	colors = append(colors, "green") // adding a new color to slice colors
	you.colors = colors              // reassigning slice colors to you.colors

	// 4.
	fmt.Printf("%v\n", me)   // {Andrei 30 [red yellow]}
	fmt.Printf("%+v\n", you) // {name:Maria age:22 colors:[pink blue green]}

	for _, colors := range me.colors {
		fmt.Printf("%s ", colors) // red yellow
	}
	fmt.Println()

	me.score.grade = 98        // update grade under score struct
	me.score.course = "Golang" // update course data

	fmt.Printf("%v\n", me)   // {Andrei 30 [red yellow] {98 Golang}}
	fmt.Printf("%+v\n", you) // {name:Maria age:22 colors:[pink blue green] score:{grade:75 course:French}}

}

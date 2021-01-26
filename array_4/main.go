package main

import (
	"fmt"
	"strings"
)

func main() {

	cities := [2]string{}
	fmt.Printf("Cities array values: %#v\n", cities) // Cities array values: [2]string{"", ""}

	grades := [3]float64{
		2.1,
		4.2,
	}
	fmt.Printf("Grades array values: %#v\n", grades) // Grades array values: [3]float64{2.1, 4.2, 0}

	taskDone := [...]bool{false, true, true, false, true}
	_ = taskDone

	fmt.Printf("TaskDone array values: %#v\n", taskDone) // TaskDone array values: [5]bool{false, true, true, false, true}

	for i := 0; i < len(cities); i++ {
		fmt.Printf("%s", cities[i])
	}
	fmt.Println() //   ==> No input == EMPTY ARRAY of STRING

	for i, value := range grades {
		fmt.Println(i, value) // INDEX with element VALUE
	}
	/*
	   0 2.1
	   1 4.2
	   2 0
	*/

	nums := []int{30, -1, -6, 90, -6}

	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 || nums[i]%2 != 0 {
			continue
		} else {
			count++
		}
	}
	fmt.Printf("There are %d positive even numbers in the array\n", count) // There are 2 positive even numbers in the array

	var count1 int

	for _, v := range nums {
		if v%2 == 0 && v > 0 {
			count1++
		}
	}

	fmt.Println("No. of positive even numbers in nums: ", count1) // No. of positive even numbers in nums:  2

	var arr1 = [3]int{}
	var arr4 = [...]int{}
	var arr2 = [4]int{1, 2, 3}
	var arr3 = [5]int{
		1,
		2,
		3,
		4,
	}
	_, _, _, _ = arr1, arr2, arr3, arr4

	fmt.Println(strings.Repeat("#", 50)) // ##################################################

	myArray := [4]float64{1.2, 5.6}

	myArray[2] = 6

	a := 10.
	myArray[0] = a

	myArray[3] = 10.10

	fmt.Println(myArray) // [10 5.6 6 10.1]
}
